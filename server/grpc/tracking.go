package grpc

import (
	context "context"

	"github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TrackingService struct {
	proto.UnimplementedTrackingServiceServer
}

func (t *TrackingService) GetsAllCreds(ctx context.Context, in *emptypb.Empty) (*proto.GetCredsResponse, error) {

	return nil, nil
}

func (t *TrackingService) GetHostInfo(ctx context.Context, in *proto.GetHostInfoRequest) (*proto.GetHostInfoResponse, error) {
	return nil, nil
}

func (t *TrackingService) GetHosts(ctx context.Context, in *emptypb.Empty) (*proto.GetHostsResponse, error) {
	return nil, nil
}
