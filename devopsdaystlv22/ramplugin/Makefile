NAME=ramdisk-plugin

docker-dev:
	docker buildx build --platform linux/arm64 --push -t cnvrg/$(NAME)-dvl:latest -f Dockerfile.dev .

sync:
	kubectl cp ./ramdiskplugin.go $(shell kubectl get pods -lapp=$(NAME) -A -ojson | jq -r '.items[] | .metadata.namespace + "/" + .metadata.name'):/opt/workdir
	kubectl cp ./Makefile $(shell kubectl get pods -lapp=$(NAME) -A -ojson | jq -r '.items[] | .metadata.namespace + "/" + .metadata.name'):/opt/workdir
	kubectl cp ./go.mod $(shell kubectl get pods -lapp=$(NAME) -A -ojson | jq -r '.items[] | .metadata.namespace + "/" + .metadata.name'):/opt/workdir

remote-debug:
	go mod tidy
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient  ramdiskplugin.go -- start

debug: sync
	 kubectl exec -n cnvrg $(NAME)  -- /bin/bash -lc "make remote-debug"