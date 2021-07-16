GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

all: compile_http-gateway compile_gw_http-gateway compile_helper-api compile_flightplan compile_flightoperation compile_action compile_flightreport compile_vehicle compile_mission cp_communication

compile_http-gateway: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto flightplan.proto flightplan_change.proto flightplan_execute.proto flightplan_assignments.proto flightoperation.proto flightreport.proto action.proto upload_mission_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightplan.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightplan_change.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightplan_execute.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightplan_assignments.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightoperation.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto flightreport.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto action.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../http-gateway/pkg/skysign_proto upload_mission_edge.proto

compile_gw_http-gateway: common.proto communication_edge.proto communication_user.proto mission.proto vehicle.proto helper_user.proto flightplan.proto flightplan_change.proto flightplan_execute.proto flightplan_assignments.proto flightoperation.proto flightreport.proto action.proto upload_mission_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto communication_edge.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto communication_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightplan.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightplan_change.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightplan_execute.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightplan_assignments.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightoperation.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto flightreport.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto action.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../http-gateway/pkg/skysign_proto upload_mission_edge.proto

compile_helper-api: helper_user.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../helper-api/pkg/skysign_proto helper_user.proto

compile_flightplan: common.proto flightplan.proto flightplan_change.proto flightplan_execute.proto flightplan_assignments.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto flightplan.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto flightplan_change.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto flightplan_execute.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto flightplan_assignments.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightplan/pkg/skysign_proto event.proto

compile_flightoperation: common.proto flightoperation.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightoperation/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightoperation/pkg/skysign_proto flightoperation.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightoperation/pkg/skysign_proto event.proto

compile_flightreport: common.proto flightreport.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightreport/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightreport/pkg/skysign_proto flightreport.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../flightreport/pkg/skysign_proto event.proto

compile_action: common.proto action.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../action/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../action/pkg/skysign_proto action.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../action/pkg/skysign_proto event.proto

compile_vehicle: common.proto vehicle.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../vehicle/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../vehicle/pkg/skysign_proto vehicle.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../vehicle/pkg/skysign_proto event.proto

compile_mission: common.proto mission.proto event.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../mission/pkg/skysign_proto common.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../mission/pkg/skysign_proto mission.proto
	protoc -I/usr/local/include -I.   -I$(GOPATH)/src   -I$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.6/third_party/googleapis   --go_out=plugins=grpc:../mission/pkg/skysign_proto event.proto

cp_communication: common.proto communication_edge.proto communication_user.proto event.proto upload_mission_edge.proto
	cp common.proto ../communication/src/main/proto
	cp communication_edge.proto ../communication/src/main/proto
	cp communication_user.proto ../communication/src/main/proto
	cp event.proto ../communication/src/main/proto
	cp upload_mission_edge.proto ../communication/src/main/proto
