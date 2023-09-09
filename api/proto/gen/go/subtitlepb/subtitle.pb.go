// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: subtitle.proto

package subtitlepb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ListVideosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor   string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"` // {"all"}
}

func (x *ListVideosRequest) Reset() {
	*x = ListVideosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVideosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVideosRequest) ProtoMessage() {}

func (x *ListVideosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVideosRequest.ProtoReflect.Descriptor instead.
func (*ListVideosRequest) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{1}
}

func (x *ListVideosRequest) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *ListVideosRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListVideosRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type ListVideosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos     []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
	NextCursor string   `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *ListVideosResponse) Reset() {
	*x = ListVideosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVideosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVideosResponse) ProtoMessage() {}

func (x *ListVideosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVideosResponse.ProtoReflect.Descriptor instead.
func (*ListVideosResponse) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{2}
}

func (x *ListVideosResponse) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

func (x *ListVideosResponse) GetNextCursor() string {
	if x != nil {
		return x.NextCursor
	}
	return ""
}

type UploadVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadVideoRequest) Reset() {
	*x = UploadVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoRequest) ProtoMessage() {}

func (x *UploadVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoRequest.ProtoReflect.Descriptor instead.
func (*UploadVideoRequest) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{3}
}

func (x *UploadVideoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UploadVideoRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UploadVideoResponse) Reset() {
	*x = UploadVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadVideoResponse) ProtoMessage() {}

func (x *UploadVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadVideoResponse.ProtoReflect.Descriptor instead.
func (*UploadVideoResponse) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{4}
}

func (x *UploadVideoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request and response types for adding subtitles
type SubtitleVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *SubtitleVideoRequest) Reset() {
	*x = SubtitleVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtitleVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtitleVideoRequest) ProtoMessage() {}

func (x *SubtitleVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtitleVideoRequest.ProtoReflect.Descriptor instead.
func (*SubtitleVideoRequest) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{5}
}

func (x *SubtitleVideoRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

type SubtitleVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SubtitleVideoResponse) Reset() {
	*x = SubtitleVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subtitle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtitleVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtitleVideoResponse) ProtoMessage() {}

func (x *SubtitleVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_subtitle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtitleVideoResponse.ProtoReflect.Descriptor instead.
func (*SubtitleVideoResponse) Descriptor() ([]byte, []int) {
	return file_subtitle_proto_rawDescGZIP(), []int{6}
}

func (x *SubtitleVideoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_subtitle_proto protoreflect.FileDescriptor

var file_subtitle_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1d, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5b,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x65, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xca, 0x03, 0x0a, 0x08,
	0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x66, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x22, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x6f, 0x6e,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73,
	0x12, 0x6c, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12,
	0x23, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x69,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x23, 0x2e,
	0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x2a, 0x07, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x7d, 0x0a, 0x0d, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x25, 0x2e, 0x64, 0x6f, 0x6e,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f,
	0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x68, 0x79, 0x6f, 0x75,
	0x6e, 0x67, 0x2f, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subtitle_proto_rawDescOnce sync.Once
	file_subtitle_proto_rawDescData = file_subtitle_proto_rawDesc
)

func file_subtitle_proto_rawDescGZIP() []byte {
	file_subtitle_proto_rawDescOnce.Do(func() {
		file_subtitle_proto_rawDescData = protoimpl.X.CompressGZIP(file_subtitle_proto_rawDescData)
	})
	return file_subtitle_proto_rawDescData
}

var file_subtitle_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_subtitle_proto_goTypes = []interface{}{
	(*Video)(nil),                 // 0: dongle.subtitle.Video
	(*ListVideosRequest)(nil),     // 1: dongle.subtitle.ListVideosRequest
	(*ListVideosResponse)(nil),    // 2: dongle.subtitle.ListVideosResponse
	(*UploadVideoRequest)(nil),    // 3: dongle.subtitle.UploadVideoRequest
	(*UploadVideoResponse)(nil),   // 4: dongle.subtitle.UploadVideoResponse
	(*SubtitleVideoRequest)(nil),  // 5: dongle.subtitle.SubtitleVideoRequest
	(*SubtitleVideoResponse)(nil), // 6: dongle.subtitle.SubtitleVideoResponse
}
var file_subtitle_proto_depIdxs = []int32{
	0, // 0: dongle.subtitle.ListVideosResponse.videos:type_name -> dongle.subtitle.Video
	1, // 1: dongle.subtitle.Subtitle.ListVideos:input_type -> dongle.subtitle.ListVideosRequest
	3, // 2: dongle.subtitle.Subtitle.UploadVideo:input_type -> dongle.subtitle.UploadVideoRequest
	3, // 3: dongle.subtitle.Subtitle.DeleteVideo:input_type -> dongle.subtitle.UploadVideoRequest
	5, // 4: dongle.subtitle.Subtitle.SubtitleVideo:input_type -> dongle.subtitle.SubtitleVideoRequest
	2, // 5: dongle.subtitle.Subtitle.ListVideos:output_type -> dongle.subtitle.ListVideosResponse
	4, // 6: dongle.subtitle.Subtitle.UploadVideo:output_type -> dongle.subtitle.UploadVideoResponse
	4, // 7: dongle.subtitle.Subtitle.DeleteVideo:output_type -> dongle.subtitle.UploadVideoResponse
	6, // 8: dongle.subtitle.Subtitle.SubtitleVideo:output_type -> dongle.subtitle.SubtitleVideoResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_subtitle_proto_init() }
func file_subtitle_proto_init() {
	if File_subtitle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subtitle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVideosRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVideosResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtitleVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_subtitle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtitleVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_subtitle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_subtitle_proto_goTypes,
		DependencyIndexes: file_subtitle_proto_depIdxs,
		MessageInfos:      file_subtitle_proto_msgTypes,
	}.Build()
	File_subtitle_proto = out.File
	file_subtitle_proto_rawDesc = nil
	file_subtitle_proto_goTypes = nil
	file_subtitle_proto_depIdxs = nil
}
