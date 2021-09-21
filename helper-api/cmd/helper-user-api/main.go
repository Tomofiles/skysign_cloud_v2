package main

import (
	"flag"
	"net"
	"os"

	"github.com/Tomofiles/skysign_cloud_v2/helper-api/pkg/api"

	cgrpc "github.com/Tomofiles/skysign_cloud_v2/skysign-common/pkg/common/adapters/grpc"

	pb "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	port *string
)

func run() error {
	port = flag.String("port", "5001", "helper api port")
	flag.Parse()
	defer glog.Flush()

	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		glog.Fatal(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(cgrpc.LogBodyInterceptor()))
	pb.RegisterHelperUserServiceServer(s, &api.Server{})

	glog.Info("start helper-api server")
	return s.Serve(lis)
}

func main() {
	if err := run(); err != nil {
		glog.Error(err)
		os.Exit(1)
	}
}
