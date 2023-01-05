GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	@go install github.com/jackc/tern@latest

.PHONY: proto
proto:
	@protoc -I . --proto_path=../../protos --micro_out=. --go_out=:. --go-grpc_out=. --go-grpc_opt=paths=source_relative --validate_out='lang=go:./' proto/userservice.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/userservicehealth.proto
	

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o userservice *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@DOCKER_BUILDKIT=1 docker build -t userservice:latest .

.PHONY: sqlc
sqlc:
	@sqlc generate -f ./postgres/sqlc.yaml