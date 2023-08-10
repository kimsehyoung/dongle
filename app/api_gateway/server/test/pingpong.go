package test

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/testpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *TestService) PingPong(ctx context.Context, empty *emptypb.Empty) (*testpb.PongResponse, error) {
	res, err := s.TestClient.PingPong(ctx, empty)
	if err != nil {
		shlog.Logf("ERROR", "Test service PingPong response err: %v", err)
		return nil, err
	}
	return &testpb.PongResponse{Message: res.Message}, nil
}
