package api

import (
	"context"

	pb "github.com/Tomofiles/skysign_cloud_v2/skysign-proto/pkg/skysign_proto"
)

// Server .
type Server struct {
	pb.UnimplementedHelperUserServiceServer
}

// GetEllipsoidHeight .
func (s *Server) GetEllipsoidHeight(ctx context.Context, in *pb.GetEllipsoidHeightRequest) (*pb.GetEllipsoidHeightResponse, error) {
	geoid := GetGeoidHeight(in.GetLatitude(), in.GetLongitude())
	elevation := GetElevation(in.GetLatitude(), in.GetLongitude())
	return &pb.GetEllipsoidHeightResponse{Height: geoid + elevation}, nil
}
