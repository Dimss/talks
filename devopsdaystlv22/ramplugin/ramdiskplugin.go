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

type FooBar struct {
	server       *grpc.Server
	socket       string
	resourceName string
}

func (b *FooBar) GetDevicePluginOptions(ctx context.Context, empty *dp.Empty) (*dp.DevicePluginOptions, error) {
	return &dp.DevicePluginOptions{}, nil
}

func (b *FooBar) ListAndWatch(empty *dp.Empty, server dp.DevicePlugin_ListAndWatchServer) error {

	for {
		devices := []*dp.Device{
			{
				ID:     "foo",
				Health: dp.Healthy,
			},
			{
				ID:     "bar",
				Health: dp.Healthy,
			},
			{
				ID:     "baz",
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

func (b *FooBar) GetPreferredAllocation(ctx context.Context, request *dp.PreferredAllocationRequest) (*dp.PreferredAllocationResponse, error) {
	return &dp.PreferredAllocationResponse{}, nil
}

func (b *FooBar) Allocate(ctx context.Context, request *dp.AllocateRequest) (*dp.AllocateResponse, error) {
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

func (b *FooBar) PreStartContainer(ctx context.Context, request *dp.PreStartContainerRequest) (*dp.PreStartContainerResponse, error) {
	return &dp.PreStartContainerResponse{}, nil
}

func (b *FooBar) Serve() error {
	_ = os.Remove(b.socket)

	dp.RegisterDevicePluginServer(b.server, b)

	sock, err := net.Listen("unix", b.socket)
	if err != nil {
		return err
	}

	go func() {
		if err := b.server.Serve(sock); err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (b *FooBar) Register() error {

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
		Endpoint:     path.Base(b.socket),
		ResourceName: b.resourceName,
		Options:      &dp.DevicePluginOptions{},
	}

	if _, err := c.Register(context.Background(), req); err != nil {
		return err
	}

	return nil
}

func main() {
	ramDisk := &FooBar{
		server:       grpc.NewServer(),
		socket:       dp.DevicePluginPath + "foobar.sock",
		resourceName: "cnvrg.io/foo-bar",
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
