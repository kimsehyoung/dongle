package video

import (
	"sync"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/videopb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VideoService struct {
	videopb.VideoServer
	VideoClient videopb.VideoClient
}

var (
	once        sync.Once
	videoClient videopb.VideoClient
)

func GetVideoClient(videoServiceAddr string) videopb.VideoClient {
	once.Do(func() {
		conn, err := grpc.Dial(videoServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logger.Errorf("can't dial to Video Serivce (%s) %v", videoServiceAddr, err)
		}
		videoClient = videopb.NewVideoClient(conn)
	})
	return videoClient
}
