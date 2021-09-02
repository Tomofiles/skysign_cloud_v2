.PHONY: dependency build-test unit-test install compile_proto #integration-test docker-up docker-down clear 

GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

dependency:
	@go mod download

build-test: dependency
	@go build -v ./...

unit-test: dependency
	@go test -v -short ./...

# integration-test: docker-up dependency
# 	@go test -v ./...

install: dependency
	@go install -v ./...

# docker-up:
# 	@docker-compose up -d --build

# docker-down:
# 	@docker-compose down

# clear: docker-down

compile_proto:
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/action.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/fleet_assignments.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/flightoperation.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/flightplan_change.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/flightplan_execute.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/flightplan.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/flightreport.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/upload_mission_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out ./pkg/skysign_proto --go-grpc_out ./pkg/skysign_proto ./proto/vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/action.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/fleet_assignments.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/flightoperation.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/flightplan_change.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/flightplan_execute.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/flightplan.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/flightreport.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/upload_mission_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out ./pkg/skysign_proto --grpc-gateway_opt logtostderr=true --grpc-gateway_opt allow_delete_body=true ./proto/vehicle.proto