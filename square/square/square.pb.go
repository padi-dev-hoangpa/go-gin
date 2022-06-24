// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: square/square.proto

package square

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

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_square_square_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_square_square_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_square_square_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Result int32 `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_square_square_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_square_square_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_square_square_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Output) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_square_square_proto protoreflect.FileDescriptor

var file_square_square_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x5c, 0x0a, 0x09, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x70, 0x63, 0x12, 0x24, 0x0a,
	0x0f, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x06, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x64, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x06, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4b,
	0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x42, 0x0b, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x68, 0x6f, 0x61, 0x6e, 0x67, 0x64, 0x7a, 0x2f, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_square_square_proto_rawDescOnce sync.Once
	file_square_square_proto_rawDescData = file_square_square_proto_rawDesc
)

func file_square_square_proto_rawDescGZIP() []byte {
	file_square_square_proto_rawDescOnce.Do(func() {
		file_square_square_proto_rawDescData = protoimpl.X.CompressGZIP(file_square_square_proto_rawDescData)
	})
	return file_square_square_proto_rawDescData
}

var file_square_square_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_square_square_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: Input
	(*Output)(nil), // 1: Output
}
var file_square_square_proto_depIdxs = []int32{
	0, // 0: SquareRpc.findSquareUnary:input_type -> Input
	0, // 1: SquareRpc.findSquareStream:input_type -> Input
	1, // 2: SquareRpc.findSquareUnary:output_type -> Output
	1, // 3: SquareRpc.findSquareStream:output_type -> Output
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_square_square_proto_init() }
func file_square_square_proto_init() {
	if File_square_square_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_square_square_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_square_square_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_square_square_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_square_square_proto_goTypes,
		DependencyIndexes: file_square_square_proto_depIdxs,
		MessageInfos:      file_square_square_proto_msgTypes,
	}.Build()
	File_square_square_proto = out.File
	file_square_square_proto_rawDesc = nil
	file_square_square_proto_goTypes = nil
	file_square_square_proto_depIdxs = nil
}
