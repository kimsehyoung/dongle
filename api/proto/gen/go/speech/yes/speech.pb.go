// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.6
// source: speech.proto

package speech

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

// ### STT ###
type RecognizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Audio    []byte `protobuf:"bytes,2,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *RecognizeRequest) Reset() {
	*x = RecognizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_speech_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeRequest) ProtoMessage() {}

func (x *RecognizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_speech_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeRequest.ProtoReflect.Descriptor instead.
func (*RecognizeRequest) Descriptor() ([]byte, []int) {
	return file_speech_proto_rawDescGZIP(), []int{0}
}

func (x *RecognizeRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *RecognizeRequest) GetAudio() []byte {
	if x != nil {
		return x.Audio
	}
	return nil
}

type RecognizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *RecognizeResponse) Reset() {
	*x = RecognizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_speech_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizeResponse) ProtoMessage() {}

func (x *RecognizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_speech_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizeResponse.ProtoReflect.Descriptor instead.
func (*RecognizeResponse) Descriptor() ([]byte, []int) {
	return file_speech_proto_rawDescGZIP(), []int{1}
}

func (x *RecognizeResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// ### TTS ###
type SynthesizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SynthesizeRequest) Reset() {
	*x = SynthesizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_speech_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeRequest) ProtoMessage() {}

func (x *SynthesizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_speech_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeRequest.ProtoReflect.Descriptor instead.
func (*SynthesizeRequest) Descriptor() ([]byte, []int) {
	return file_speech_proto_rawDescGZIP(), []int{2}
}

func (x *SynthesizeRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *SynthesizeRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SynthesizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audio []byte `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio,omitempty"`
}

func (x *SynthesizeResponse) Reset() {
	*x = SynthesizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_speech_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SynthesizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SynthesizeResponse) ProtoMessage() {}

func (x *SynthesizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_speech_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SynthesizeResponse.ProtoReflect.Descriptor instead.
func (*SynthesizeResponse) Descriptor() ([]byte, []int) {
	return file_speech_proto_rawDescGZIP(), []int{3}
}

func (x *SynthesizeResponse) GetAudio() []byte {
	if x != nil {
		return x.Audio
	}
	return nil
}

var File_speech_proto protoreflect.FileDescriptor

var file_speech_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x22, 0x27, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x53, 0x79,
	0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0x2a, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x32, 0xe6, 0x01, 0x0a, 0x06,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x12, 0x6c, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x22, 0x11, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x7a, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x2e, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x22, 0x10, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65,
	0x73, 0x69, 0x7a, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6d, 0x73, 0x65, 0x68, 0x79, 0x6f, 0x75, 0x6e, 0x67, 0x2f, 0x64,
	0x6f, 0x6e, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_speech_proto_rawDescOnce sync.Once
	file_speech_proto_rawDescData = file_speech_proto_rawDesc
)

func file_speech_proto_rawDescGZIP() []byte {
	file_speech_proto_rawDescOnce.Do(func() {
		file_speech_proto_rawDescData = protoimpl.X.CompressGZIP(file_speech_proto_rawDescData)
	})
	return file_speech_proto_rawDescData
}

var file_speech_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_speech_proto_goTypes = []interface{}{
	(*RecognizeRequest)(nil),   // 0: dongle.speech.RecognizeRequest
	(*RecognizeResponse)(nil),  // 1: dongle.speech.RecognizeResponse
	(*SynthesizeRequest)(nil),  // 2: dongle.speech.SynthesizeRequest
	(*SynthesizeResponse)(nil), // 3: dongle.speech.SynthesizeResponse
}
var file_speech_proto_depIdxs = []int32{
	0, // 0: dongle.speech.Speech.Recognize:input_type -> dongle.speech.RecognizeRequest
	2, // 1: dongle.speech.Speech.Synthesize:input_type -> dongle.speech.SynthesizeRequest
	1, // 2: dongle.speech.Speech.Recognize:output_type -> dongle.speech.RecognizeResponse
	3, // 3: dongle.speech.Speech.Synthesize:output_type -> dongle.speech.SynthesizeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_speech_proto_init() }
func file_speech_proto_init() {
	if File_speech_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_speech_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeRequest); i {
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
		file_speech_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizeResponse); i {
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
		file_speech_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeRequest); i {
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
		file_speech_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SynthesizeResponse); i {
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
			RawDescriptor: file_speech_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_speech_proto_goTypes,
		DependencyIndexes: file_speech_proto_depIdxs,
		MessageInfos:      file_speech_proto_msgTypes,
	}.Build()
	File_speech_proto = out.File
	file_speech_proto_rawDesc = nil
	file_speech_proto_goTypes = nil
	file_speech_proto_depIdxs = nil
}