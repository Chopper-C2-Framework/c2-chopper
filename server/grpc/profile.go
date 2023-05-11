package grpc

import (
	context "context"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
)

type ProfileService struct {
	proto.UnimplementedProfileServiceServer
}

func (s *ProfileService) CreateProfile(context.Context, *proto.CreateProfileRequest) (*proto.CreateProfileResponse, error) {
	return nil, nil
}
func (s *ProfileService) UpdateProfile(context.Context, *proto.UpdateProfileRequest) (*proto.UpdateProfileResponse, error) {
	return nil, nil
}
