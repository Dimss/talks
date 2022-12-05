package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	dp "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"os"
	"os/signal"
	"path"
	"syscall"

	"net"
	"time"
)

type RamPlugin struct {
	server       *grpc.Server
	socket       string
	resourceName string
}

func (p *RamPlugin) GetDevicePluginOptions(ctx context.Context, empty *dp.Empty) (*dp.DevicePluginOptions, error) {
	return &dp.DevicePluginOptions{}, nil
}

func (p *RamPlugin) ListAndWatch(empty *dp.Empty, server dp.DevicePlugin_ListAndWatchServer) error {

	for {
		devices := []*dp.Device{
			{
				ID:     "/mnt/disk1",
				Health: dp.Healthy,
			},
			{
				ID:     "/mnt/disk2",
				Health: dp.Healthy,
			},
			{
				ID:     "/mnt/disk3",
				Health: dp.Healthy,
			},
		}
		response := &dp.ListAndWatchResponse{Devices: devices}
		if err := server.Send(response); err != nil {
			log.Error(err)
		}
		time.Sleep(300 * time.Second)
	}
}

func (p *RamPlugin) GetPreferredAllocation(ctx context.Context, request *dp.PreferredAllocationRequest) (*dp.PreferredAllocationResponse, error) {
	return &dp.PreferredAllocationResponse{}, nil
}

func (p *RamPlugin) Allocate(ctx context.Context, request *dp.AllocateRequest) (*dp.AllocateResponse, error) {
	allocResponse := &dp.AllocateResponse{}
	for _, req := range request.ContainerRequests {
		containerResponse := &dp.ContainerAllocateResponse{}
		for _, ramDisk := range req.DevicesIDs {
			containerResponse.Mounts = append(containerResponse.Mounts, &dp.Mount{
				ContainerPath: ramDisk,
				HostPath:      ramDisk,
			})
		}
		allocResponse.ContainerResponses = append(allocResponse.ContainerResponses, containerResponse)
	}

	return allocResponse, nil
}

func (p *RamPlugin) PreStartContainer(ctx context.Context, request *dp.PreStartContainerRequest) (*dp.PreStartContainerResponse, error) {
	return &dp.PreStartContainerResponse{}, nil
}

func (p *RamPlugin) Serve() error {
	_ = os.Remove(p.socket)

	dp.RegisterDevicePluginServer(p.server, p)

	sock, err := net.Listen("unix", p.socket)
	if err != nil {
		return err
	}

	go func() {
		if err := p.server.Serve(sock); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (p *RamPlugin) Register() error {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return net.DialTimeout("unix", s, 5*time.Second)
		}),
	}

	conn, err := grpc.Dial(dp.KubeletSocket, opts...)

	if err != nil {
		return err
	}

	c := dp.NewRegistrationClient(conn)

	req := &dp.RegisterRequest{
		Version:      dp.Version,
		Endpoint:     path.Base(p.socket),
		ResourceName: p.resourceName,
		Options:      &dp.DevicePluginOptions{},
	}

	if _, err := c.Register(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func main() {
	ramDisk := &RamPlugin{
		server:       grpc.NewServer(),
		socket:       dp.DevicePluginPath + "ramdisk.sock",
		resourceName: "cnvrg.io/ramdisk",
	}
	if err := ramDisk.Serve(); err != nil {
		log.Fatal(err)
	}

	if err := ramDisk.Register(); err != nil {
		log.Fatal(err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)
	_ = <-sigCh
	os.Exit(0)

}
