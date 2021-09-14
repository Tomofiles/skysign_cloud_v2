package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// NewGrpcClientConnection .
func NewGrpcClientConnectionWithBlockAndTimeout(url string) (*grpc.ClientConn, error) {
	return grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
}
