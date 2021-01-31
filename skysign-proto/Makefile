GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

compile_http_gateway: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto helper_user.proto

compile_gw_http_gateway: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto helper_user.proto

compile_helper-api: helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../helper-api/pkg/skysign_proto helper_user.proto

cp_communication: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto event.proto
	cp common.proto ../communication/src/main/proto
	cp communication_edge.proto ../communication/src/main/proto
	cp communication_user.proto ../communication/src/main/proto
	cp mission.proto ../communication/src/main/proto
	cp vehicle.proto ../communication/src/main/proto
	cp helper_user.proto ../communication/src/main/proto
	cp event.proto ../communication/src/main/proto

cp_mission: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto event.proto
	cp common.proto ../mission/src/main/proto
	cp communication_edge.proto ../mission/src/main/proto
	cp communication_user.proto ../mission/src/main/proto
	cp mission.proto ../mission/src/main/proto
	cp vehicle.proto ../mission/src/main/proto
	cp helper_user.proto ../mission/src/main/proto
	cp event.proto ../mission/src/main/proto
