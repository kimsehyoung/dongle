package speech

import (
	"sync"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/speechpb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SpeechService struct {
	speechpb.SpeechServer
	SpeechClient speechpb.SpeechClient
}

var (
	once         sync.Once
	speechClient speechpb.SpeechClient
)

func GetSpeechClient(speechServiceAddr string) speechpb.SpeechClient {
	once.Do(func() {
		conn, err := grpc.Dial(speechServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Errorf("can't dial to Speech Serivce (%s) %v", speechServiceAddr, err)
		}
		speechClient = speechpb.NewSpeechClient(conn)
	})
	return speechClient
}
