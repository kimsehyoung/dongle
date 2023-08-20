package google

import (
	"context"
	"sync"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kimsehyoung/logger"
)

type GoogleClient struct {
	*texttospeech.Client
}

var (
	once         sync.Once
	googleClient *GoogleClient
)

func GetGoogleClient(ctx context.Context) *GoogleClient {
	once.Do(func() {
		client, err := texttospeech.NewClient(ctx)
		if err != nil {
			logger.Errorf("Faield to intialize google client: %v", err)
			return
		}
		googleClient = &GoogleClient{client}
	})
	return googleClient
}

func (c *GoogleClient) TextToSpeech(ctx context.Context, body TtsRequest) (*TtsResponse, error) {
	googleReq := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: body.Text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: body.Language,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_LINEAR16,
		},
	}
	// Request SynthesizeSpeech to Google Cloud
	resp, err := c.SynthesizeSpeech(ctx, &googleReq)
	if err != nil {
		logger.Errorf("Request Failed Test service PingPong response err: %v", err)
		return nil, err
	}
	ttsReponse := TtsResponse{
		Audio: resp.AudioContent,
	}

	return &ttsReponse, nil
}
