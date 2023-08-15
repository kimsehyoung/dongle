package speech

import (
	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/gopackages/shlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SpeechService struct {
	speechpb.SpeechServer
	SpeechClient speechpb.SpeechClient
}

func GetSpeechClient(speechServiceAddr string) speechpb.SpeechClient {
	conn, err := grpc.Dial(speechServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		shlog.Logf("ERROR", "can't dial to Speech Serivce (%s) %v", speechServiceAddr, err)
	}
	return speechpb.NewSpeechClient(conn)
}
