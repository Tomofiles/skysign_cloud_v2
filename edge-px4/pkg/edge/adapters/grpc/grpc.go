package grpc

import "google.golang.org/grpc"

// NewGrpcClientConnection .
func NewGrpcClientConnectionWithBlock(url string) (*grpc.ClientConn, error) {
	return grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
}
