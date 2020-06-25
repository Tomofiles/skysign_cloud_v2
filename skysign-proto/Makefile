GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

http_gateway/pkg/skysign_proto/communication.pb.go: communication.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http_gateway/pkg/skysign_proto communication.proto

http_gateway/pkg/skysign_proto/communication.pb.gw.go: communication.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http_gateway/pkg/skysign_proto communication.proto

communication/src/main/proto/communication.proto: communication.proto
	cp communication.proto ../communication/src/main/proto