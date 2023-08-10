package speech

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/app/api_gateway/server"
	"github.com/kimsehyoung/gopackages/shlog"
)

func (s *server.ApiGatewayServer) Synthesize(ctx context.Context, req *speechpb.SynthesizeRequest) (*speechpb.SynthesizeResponse, error) {
	shlog.Logf("INFO", "Start")
	return nil, nil
}
