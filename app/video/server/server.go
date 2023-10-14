package server

import (
	"context"
	"net"

	"github.com/kimsehyoung/dongle/api/proto/gen/go/videopb"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen"
	"github.com/kimsehyoung/dongle/app/video/server/aws"
	"github.com/kimsehyoung/dongle/app/video/server/videodb"
	"github.com/kimsehyoung/logger"
	"google.golang.org/grpc"
)

type ServiceInfo struct {
	VideoServiceAddr string
	DatabaseInfo     videodb.DatabaseInfo
}

type videoService struct {
	videopb.VideoServer
	videodbClient *videodbgen.Client
	s3Client      *aws.S3Client
}

func StartGrpcServer(serviceInfo ServiceInfo) {
	// Register gRPC Server
	grpcServer := grpc.NewServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	videopb.RegisterVideoServer(grpcServer, &videoService{
		videodbClient: videodb.GetDbClient(serviceInfo.DatabaseInfo),
		s3Client:      aws.GetS3Client(ctx),
	})

	// Start gRPC server
	lis, err := net.Listen("tcp", serviceInfo.VideoServiceAddr)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	logger.Infof("gRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
