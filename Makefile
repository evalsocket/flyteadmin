export REPOSITORY=flyteadmin
include boilerplate/lyft/docker_build/Makefile
include boilerplate/lyft/golang_test_targets/Makefile
include boilerplate/lyft/end2end/Makefile

.PHONY: update_boilerplate
update_boilerplate:
	@boilerplate/update.sh

.PHONY: integration
integration:
	CGO_ENABLED=0 GOFLAGS="-count=1" go test -v -tags=integration ./tests/...

.PHONY: k8s_integration
k8s_integration:
	@script/integration/launch.sh

.PHONY: k8s_integration_execute
k8s_integration_execute:
	@script/integration/execute.sh

.PHONY: compile
compile:
	go build -o flyteadmin ./cmd/ && mv ./flyteadmin ${GOPATH}/bin

.PHONY: linux_compile
linux_compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /artifacts/flyteadmin ./cmd/

.PHONY: server
server:
	go run cmd/main.go --logtostderr --server.kube-config ~/.kube/config  --config flyteadmin_config.yaml serve

.PHONY: migrate
migrate:
	go run cmd/main.go --logtostderr --server.kube-config ~/.kube/config  --config flyteadmin_config.yaml migrate run

.PHONY: seed_projects
seed_projects:
	go run cmd/main.go --logtostderr --server.kube-config ~/.kube/config  --config flyteadmin_config.yaml migrate seed-projects project admintests flytekit

all: compile

generate: download_tooling
	@go generate ./...
