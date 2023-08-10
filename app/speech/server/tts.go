package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func (s *speechService) Synthesize(ctx context.Context, req *speechpb.SynthesizeRequest) (*speechpb.SynthesizeResponse, error) {
	// client, err := texttospeech.NewClient(ctx)
	// if err != nil {
	// 	return err
	// }
	// defer client.Close()

	// text := req.Text

	googleReq := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: string(text)},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: req.Language,
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &googleReq)
	// if err != nil {
	// 	shlog.Logf("ERROR", "Request Failed Test service PingPong response err: %v", err)
	// 	return nil, err
	// }
	// outputFile := "output.mp3"
	// err = ioutil.WriteFile(outputFile, resp.AudioContent, 0644)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Fprintf(w, "Audio content written to file: %v\n", outputFile)
	return nil, nil

}
