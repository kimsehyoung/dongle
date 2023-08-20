package server

import (
	"context"
	"encoding/base64"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/dongle/app/speech/server/google"
	"github.com/kimsehyoung/dongle/app/speech/server/whisper"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *speechService) Recognize(ctx context.Context, req *speechpb.RecognizeRequest) (*speechpb.RecognizeResponse, error) {

	request := whisper.SttRequest{
		Audio:    base64.StdEncoding.EncodeToString(req.Audio),
		Language: req.Language,
	}
	speechResp, err := s.whisperClient.SpeechToText(whisper.TRANSCRIPTION_PATH, request)
	if err != nil {
		logger.Errorf("Failed to request whisper SpeechToText: %v", err)
		return nil, err
	}
	return &speechpb.RecognizeResponse{Text: speechResp.Text}, nil
}

func (s *speechService) Synthesize(ctx context.Context, req *speechpb.SynthesizeRequest) (*speechpb.SynthesizeResponse, error) {

	lang := google.ConvertLang(req.Language)
	if lang == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Unsupported Language: %s", req.Language)
	}

	request := google.TtsRequest{
		Text:     req.Text,
		Language: lang,
	}

	ttsResp, err := s.googleClient.TextToSpeech(ctx, request)
	if err != nil {
		logger.Errorf("Failed to request google TextToSpeech: %v", err)
		return nil, err
	}
	return &speechpb.SynthesizeResponse{Audio: []byte(ttsResp.Audio)}, nil
}
