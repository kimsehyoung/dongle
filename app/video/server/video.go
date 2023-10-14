package server

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/videopb"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	bucket = "dongle-video"
)

func (s *videoService) UploadVideo(ctx context.Context, req *videopb.UploadVideoRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *videoService) DeleteVideo(ctx context.Context, req *videopb.DeleteVideoRequest) (*emptypb.Empty, error) {

	return nil, nil
}

func (s *videoService) ListVideos(ctx context.Context, req *videopb.ListVideosRequest) (*videopb.ListVideosResponse, error) {

	return nil, nil

}

func (s *videoService) SubtitleVideo(ctx context.Context, req *videopb.SubtitleVideoRequest) (*emptypb.Empty, error) {

	return nil, nil
}
