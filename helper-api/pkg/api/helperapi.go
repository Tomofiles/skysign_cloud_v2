package api

import (
	"context"

	pb "github.com/Tomofiles/skysign_cloud/helper-api/pkg/skysign_proto"
)

// Server .
type Server struct{}

// GetEllipsoidHeight .
func (s *Server) GetEllipsoidHeight(ctx context.Context, in *pb.GetEllipsoidHeightRequest) (*pb.GetEllipsoidHeightResponse, error) {
	geoid := GetGeoidHeight(in.GetLatitude(), in.GetLongitude())
	elevation := GetElevation(in.GetLatitude(), in.GetLongitude())
	return &pb.GetEllipsoidHeightResponse{Height: geoid + elevation}, nil
}
