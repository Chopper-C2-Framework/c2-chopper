package server_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/chopper-c2-framework/c2-chopper/grpc/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func gRPCPluginService_Test(t *testing.T, conn *grpc.ClientConn) {
	client := pb.NewPluginServiceClient(conn)

	t.Run("ListPlugins", func(subTest *testing.T) {
		gRPCListPlugins_Test(subTest, client)
	})
	t.Run("ListLoadedPlugins", func(subTest *testing.T) {
		gRPCListLoadedPlugins_Test(subTest, client)
	})
	t.Run("LoadPlugin", func(subTest *testing.T) {
		gRPCLoadPlugin_Test(subTest, client)
	})
}

func gRPCListPlugins_Test(t *testing.T, client pb.PluginServiceClient) {
	ctx := context.Background()

	type expectation struct {
		out *pb.ListPluginsResponse
		err error
	}

	tests := map[string]struct {
		in       *emptypb.Empty
		expected expectation
	}{
		"ListPlugins_Must_Success": {
			in: &emptypb.Empty{},
			expected: expectation{
				out: &pb.ListPluginsResponse{
					Names:   make([]string, 0),
					Success: true,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.ListPlugins(ctx, tt.in)
			if err != nil {
				if tt.expected.err == nil {
					t.Errorf("Err -> \nWant: no error\nGot: %q\n", err)
				} else if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if out.Success != tt.expected.out.Success {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}

func gRPCListLoadedPlugins_Test(t *testing.T, client pb.PluginServiceClient) {
	ctx := context.Background()

	type expectation struct {
		out *pb.ListLoadedPluginsResponse
		err error
	}

	tests := map[string]struct {
		in       *emptypb.Empty
		expected expectation
	}{
		"ListLoadedPlugins_Must_Success": {
			in: &emptypb.Empty{},
			expected: expectation{
				out: &pb.ListLoadedPluginsResponse{
					Names:   make([]string, 0),
					Success: true,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.ListLoadedPlugins(ctx, tt.in)
			if err != nil {
				if tt.expected.err == nil {
					t.Errorf("Err -> \nWant: no error\nGot: %q\n", err)
				} else if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if out.Success != tt.expected.out.Success {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}
			fmt.Println("LAST HERE")
		})
	}
}

func gRPCLoadPlugin_Test(t *testing.T, client pb.PluginServiceClient) {
	ctx := context.Background()

	type expectation struct {
		out *pb.LoadPluginResponse
		err error
	}

	tests := map[string]struct {
		in       *pb.LoadPluginRequest
		expected expectation
	}{
		"LoadPlugin_Invalid_Path": {
			in: &pb.LoadPluginRequest{
				FileName: "Random Invalid File Name",
			},
			expected: expectation{
				out: &pb.LoadPluginResponse{},
				err: errors.New("rpc error: code = Unknown desc = plugin not found"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.LoadPlugin(ctx, tt.in)
			if err != nil {
				if tt.expected.err == nil {
					t.Errorf("Err -> \nWant: no error\nGot: %q\n", err)
				} else if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if out.GetData() == nil {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}
