// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: grpc/chat/chat.proto

package chat

import (
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid     uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Time    uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_grpc_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Message) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_grpc_chat_chat_proto_rawDescGZIP(), []int{1}
}

var File_grpc_chat_chat_proto protoreflect.FileDescriptor

var file_grpc_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x22,
	0x49, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f,
	0x69, 0x64, 0x32, 0x3c, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x28, 0x01,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_chat_chat_proto_rawDescOnce sync.Once
	file_grpc_chat_chat_proto_rawDescData = file_grpc_chat_chat_proto_rawDesc
)

func file_grpc_chat_chat_proto_rawDescGZIP() []byte {
	file_grpc_chat_chat_proto_rawDescOnce.Do(func() {
		file_grpc_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_chat_chat_proto_rawDescData)
	})
	return file_grpc_chat_chat_proto_rawDescData
}

var file_grpc_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_chat_chat_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: connect.Message
	(*Void)(nil),    // 1: connect.Void
}
var file_grpc_chat_chat_proto_depIdxs = []int32{
	0, // 0: connect.ChatService.OpenChat:input_type -> connect.Message
	1, // 1: connect.ChatService.OpenChat:output_type -> connect.Void
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_chat_chat_proto_init() }
func file_grpc_chat_chat_proto_init() {
	if File_grpc_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_grpc_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_grpc_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_chat_chat_proto_goTypes,
		DependencyIndexes: file_grpc_chat_chat_proto_depIdxs,
		MessageInfos:      file_grpc_chat_chat_proto_msgTypes,
	}.Build()
	File_grpc_chat_chat_proto = out.File
	file_grpc_chat_chat_proto_rawDesc = nil
	file_grpc_chat_chat_proto_goTypes = nil
	file_grpc_chat_chat_proto_depIdxs = nil
}