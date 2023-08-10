package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *testService) PingPong(context.Context, *emptypb.Empty) (*testpb.PongResponse, error) {
	shlog.Log("INFO", "Ping ")
	return &testpb.PongResponse{Message: "Pong"}, nil
}
