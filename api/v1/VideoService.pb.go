// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.27.0--dev
// source: VideoService.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_VideoService_proto protoreflect.FileDescriptor

var file_VideoService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x1a, 0x13, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x46, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x12, 0x39, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x13, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_VideoService_proto_goTypes = []interface{}{
	(*VideoRequest)(nil),  // 0: video.VideoRequest
	(*VideoResponse)(nil), // 1: video.VideoResponse
}
var file_VideoService_proto_depIdxs = []int32{
	0, // 0: video.Converter.ConvertVideo:input_type -> video.VideoRequest
	1, // 1: video.Converter.ConvertVideo:output_type -> video.VideoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_VideoService_proto_init() }
func file_VideoService_proto_init() {
	if File_VideoService_proto != nil {
		return
	}
	file_VideoResponse_proto_init()
	file_VideoRequest_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_VideoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_VideoService_proto_goTypes,
		DependencyIndexes: file_VideoService_proto_depIdxs,
	}.Build()
	File_VideoService_proto = out.File
	file_VideoService_proto_rawDesc = nil
	file_VideoService_proto_goTypes = nil
	file_VideoService_proto_depIdxs = nil
}
