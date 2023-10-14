// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: video.proto

package videopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Video_UploadVideo_FullMethodName   = "/dongle.video.Video/UploadVideo"
	Video_DeleteVideo_FullMethodName   = "/dongle.video.Video/DeleteVideo"
	Video_ListVideos_FullMethodName    = "/dongle.video.Video/ListVideos"
	Video_SubtitleVideo_FullMethodName = "/dongle.video.Video/SubtitleVideo"
)

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListVideos(ctx context.Context, in *ListVideosRequest, opts ...grpc.CallOption) (*ListVideosResponse, error)
	SubtitleVideo(ctx context.Context, in *SubtitleVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) UploadVideo(ctx context.Context, in *UploadVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_UploadVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_DeleteVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) ListVideos(ctx context.Context, in *ListVideosRequest, opts ...grpc.CallOption) (*ListVideosResponse, error) {
	out := new(ListVideosResponse)
	err := c.cc.Invoke(ctx, Video_ListVideos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) SubtitleVideo(ctx context.Context, in *SubtitleVideoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Video_SubtitleVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	UploadVideo(context.Context, *UploadVideoRequest) (*emptypb.Empty, error)
	DeleteVideo(context.Context, *DeleteVideoRequest) (*emptypb.Empty, error)
	ListVideos(context.Context, *ListVideosRequest) (*ListVideosResponse, error)
	SubtitleVideo(context.Context, *SubtitleVideoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) UploadVideo(context.Context, *UploadVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoServer) ListVideos(context.Context, *ListVideosRequest) (*ListVideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideos not implemented")
}
func (UnimplementedVideoServer) SubtitleVideo(context.Context, *SubtitleVideoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubtitleVideo not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).UploadVideo(ctx, req.(*UploadVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_DeleteVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_ListVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVideosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).ListVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_ListVideos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).ListVideos(ctx, req.(*ListVideosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_SubtitleVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtitleVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).SubtitleVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Video_SubtitleVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).SubtitleVideo(ctx, req.(*SubtitleVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dongle.video.Video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadVideo",
			Handler:    _Video_UploadVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _Video_DeleteVideo_Handler,
		},
		{
			MethodName: "ListVideos",
			Handler:    _Video_ListVideos_Handler,
		},
		{
			MethodName: "SubtitleVideo",
			Handler:    _Video_SubtitleVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
