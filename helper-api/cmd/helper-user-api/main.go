package main

import (
	"flag"
	"net"

	"helper-api/pkg/api"

	pb "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port *string
)

func main() {
	port = flag.String("port", "5001", "helper api port")
	flag.Parse()
	defer glog.Flush()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		glog.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelperUserServiceServer(s, &api.Server{})
	if err := s.Serve(lis); err != nil {
		glog.Fatal(err)
	}
}
