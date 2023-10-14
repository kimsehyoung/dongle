package video

import (
	"context"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/videopb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *VideoService) UploadVideo(ctx context.Context, req *videopb.UploadVideoRequest) (*emptypb.Empty, error) {
	logger.Infof("Start")
	resp, err := s.VideoClient.UploadVideo(ctx, req)
	if err != nil {
		logger.Errorf("Failed to respond from Subtitle Service (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *VideoService) DeleteVideo(ctx context.Context, req *videopb.DeleteVideoRequest) (*emptypb.Empty, error) {
	logger.Infof("Start")
	resp, err := s.VideoClient.DeleteVideo(ctx, req)
	if err != nil {
		logger.Errorf("Failed to respond from Subtitle Service (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *VideoService) ListVideos(ctx context.Context, req *videopb.ListVideosRequest) (*videopb.ListVideosResponse, error) {

	resp, err := s.VideoClient.ListVideos(ctx, req)
	if err != nil {
		logger.Errorf("Failed to respond from Subtitle Service (%v)", err)
		return nil, err
	}
	return resp, nil
}

func (s *VideoService) SubtitleVideo(ctx context.Context, req *videopb.SubtitleVideoRequest) (*emptypb.Empty, error) {
	logger.Infof("Start")
	resp, err := s.VideoClient.SubtitleVideo(ctx, req)
	if err != nil {
		logger.Errorf("Failed to respond from Subtitle Service (%v)", err)
		return nil, err
	}
	return resp, nil
}
