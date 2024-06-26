// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: csvparser/csvparser.proto

package grpc_test_drive

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

type CSVParserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName string `protobuf:"bytes,1,opt,name=columnName,proto3" json:"columnName,omitempty"`
}

func (x *CSVParserRequest) Reset() {
	*x = CSVParserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csvparser_csvparser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSVParserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSVParserRequest) ProtoMessage() {}

func (x *CSVParserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_csvparser_csvparser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSVParserRequest.ProtoReflect.Descriptor instead.
func (*CSVParserRequest) Descriptor() ([]byte, []int) {
	return file_csvparser_csvparser_proto_rawDescGZIP(), []int{0}
}

func (x *CSVParserRequest) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

type CSVParserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowCount int64 `protobuf:"varint,1,opt,name=rowCount,proto3" json:"rowCount,omitempty"`
}

func (x *CSVParserResponse) Reset() {
	*x = CSVParserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csvparser_csvparser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSVParserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSVParserResponse) ProtoMessage() {}

func (x *CSVParserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_csvparser_csvparser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSVParserResponse.ProtoReflect.Descriptor instead.
func (*CSVParserResponse) Descriptor() ([]byte, []int) {
	return file_csvparser_csvparser_proto_rawDescGZIP(), []int{1}
}

func (x *CSVParserResponse) GetRowCount() int64 {
	if x != nil {
		return x.RowCount
	}
	return 0
}

var File_csvparser_csvparser_proto protoreflect.FileDescriptor

var file_csvparser_csvparser_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x73, 0x76, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x73, 0x76, 0x70,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x10, 0x43,
	0x53, 0x56, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x2f, 0x0a, 0x11, 0x43, 0x53, 0x56, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x41, 0x0a, 0x09, 0x43, 0x53, 0x56, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a,
	0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x11, 0x2e, 0x43, 0x53, 0x56,
	0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x43, 0x53, 0x56, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x76, 0x67, 0x65, 0x6e, 0x69, 0x79, 0x61, 0x72, 0x62, 0x61, 0x74, 0x6f, 0x76,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_csvparser_csvparser_proto_rawDescOnce sync.Once
	file_csvparser_csvparser_proto_rawDescData = file_csvparser_csvparser_proto_rawDesc
)

func file_csvparser_csvparser_proto_rawDescGZIP() []byte {
	file_csvparser_csvparser_proto_rawDescOnce.Do(func() {
		file_csvparser_csvparser_proto_rawDescData = protoimpl.X.CompressGZIP(file_csvparser_csvparser_proto_rawDescData)
	})
	return file_csvparser_csvparser_proto_rawDescData
}

var file_csvparser_csvparser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_csvparser_csvparser_proto_goTypes = []interface{}{
	(*CSVParserRequest)(nil),  // 0: CSVParserRequest
	(*CSVParserResponse)(nil), // 1: CSVParserResponse
}
var file_csvparser_csvparser_proto_depIdxs = []int32{
	0, // 0: CSVParser.CountRows:input_type -> CSVParserRequest
	1, // 1: CSVParser.CountRows:output_type -> CSVParserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_csvparser_csvparser_proto_init() }
func file_csvparser_csvparser_proto_init() {
	if File_csvparser_csvparser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csvparser_csvparser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSVParserRequest); i {
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
		file_csvparser_csvparser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSVParserResponse); i {
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
			RawDescriptor: file_csvparser_csvparser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csvparser_csvparser_proto_goTypes,
		DependencyIndexes: file_csvparser_csvparser_proto_depIdxs,
		MessageInfos:      file_csvparser_csvparser_proto_msgTypes,
	}.Build()
	File_csvparser_csvparser_proto = out.File
	file_csvparser_csvparser_proto_rawDesc = nil
	file_csvparser_csvparser_proto_goTypes = nil
	file_csvparser_csvparser_proto_depIdxs = nil
}
