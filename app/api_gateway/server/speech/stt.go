package speech

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/gopackages/shlog"
)

func (s *server.ApiGatewayServer) Recognize(ctx context.Context, req *speechpb.RecognizeRequest) (*speechpb.RecognizeResponse, error) {
	shlog.Logf("INFO", "Start")
	return nil, nil
}
