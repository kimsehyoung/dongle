package speech

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/gopackages/shlog"
)

func (s *SpeechService) Recognize(ctx context.Context, req *speechpb.RecognizeRequest) (*speechpb.RecognizeResponse, error) {

	resp, err := s.SpeechClient.Recognize(ctx, req)
	if err != nil {
		shlog.Logf("ERROR", "Failed to respond from Speech Service (%v)", err)
		return nil, err
	}
	return &speechpb.RecognizeResponse{Text: resp.Text}, nil
}

func (s *SpeechService) Synthesize(ctx context.Context, req *speechpb.SynthesizeRequest) (*speechpb.SynthesizeResponse, error) {
	shlog.Logf("INFO", "Start")
	resp, err := s.SpeechClient.Synthesize(ctx, req)
	if err != nil {
		shlog.Logf("ERROR", "Failed to respond from Speech Service (%v)", err)
		return nil, err
	}
	return &speechpb.SynthesizeResponse{Audio: resp.Audio}, nil
}
