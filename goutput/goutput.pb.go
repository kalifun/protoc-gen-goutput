// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: goutput.proto

package goutput

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_goutput_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         71001,
		Name:          "goutput.method",
		Tag:           "bytes,71001,opt,name=method",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         71002,
		Name:          "goutput.info",
		Tag:           "bytes,71002,opt,name=info",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         71003,
		Name:          "goutput.path",
		Tag:           "bytes,71003,opt,name=path",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*int64)(nil),
		Field:         71004,
		Name:          "goutput.limit",
		Tag:           "varint,71004,opt,name=limit",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         71005,
		Name:          "goutput.version",
		Tag:           "bytes,71005,opt,name=version",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         72001,
		Name:          "goutput.required",
		Tag:           "varint,72001,opt,name=required",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         72002,
		Name:          "goutput.desc",
		Tag:           "bytes,72002,opt,name=desc",
		Filename:      "goutput.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         72003,
		Name:          "goutput.jsontag",
		Tag:           "bytes,72003,opt,name=jsontag",
		Filename:      "goutput.proto",
	},
}

// Extension fields to descriptorpb.ServiceOptions.
var (
	// method GET POST
	//
	// optional string method = 71001;
	E_Method = &file_goutput_proto_extTypes[0]
	// description method
	//
	// optional string info = 71002;
	E_Info = &file_goutput_proto_extTypes[1]
	// api path
	//
	// optional string path = 71003;
	E_Path = &file_goutput_proto_extTypes[2]
	// speed limiter
	//
	// optional int64 limit = 71004;
	E_Limit = &file_goutput_proto_extTypes[3]
	// api version
	//
	// optional string version = 71005;
	E_Version = &file_goutput_proto_extTypes[4]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// required parameters
	//
	// optional bool required = 72001;
	E_Required = &file_goutput_proto_extTypes[5]
	// field Description
	//
	// optional string desc = 72002;
	E_Desc = &file_goutput_proto_extTypes[6]
	// json tag
	//
	// optional string jsontag = 72003;
	E_Jsontag = &file_goutput_proto_extTypes[7]
)

var File_goutput_proto protoreflect.FileDescriptor

var file_goutput_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3c, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0xaa, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x38, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xda, 0xaa, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88,
	0x01, 0x01, 0x3a, 0x38, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0xaa, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x3a, 0x3a, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0xaa, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x3a, 0x3e, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdd, 0xaa, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x3a, 0x3e, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xc1, 0xb2, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x88, 0x01, 0x01, 0x3a, 0x36, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xc2, 0xb2, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x88, 0x01, 0x01,
	0x3a, 0x3c, 0x0a, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc3, 0xb2, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x74, 0x61, 0x67, 0x88, 0x01, 0x01, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x6c,
	0x69, 0x66, 0x75, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_goutput_proto_goTypes = []interface{}{
	(*descriptorpb.ServiceOptions)(nil), // 0: google.protobuf.ServiceOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_goutput_proto_depIdxs = []int32{
	0, // 0: goutput.method:extendee -> google.protobuf.ServiceOptions
	0, // 1: goutput.info:extendee -> google.protobuf.ServiceOptions
	0, // 2: goutput.path:extendee -> google.protobuf.ServiceOptions
	0, // 3: goutput.limit:extendee -> google.protobuf.ServiceOptions
	0, // 4: goutput.version:extendee -> google.protobuf.ServiceOptions
	1, // 5: goutput.required:extendee -> google.protobuf.FieldOptions
	1, // 6: goutput.desc:extendee -> google.protobuf.FieldOptions
	1, // 7: goutput.jsontag:extendee -> google.protobuf.FieldOptions
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	0, // [0:8] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goutput_proto_init() }
func file_goutput_proto_init() {
	if File_goutput_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goutput_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 8,
			NumServices:   0,
		},
		GoTypes:           file_goutput_proto_goTypes,
		DependencyIndexes: file_goutput_proto_depIdxs,
		ExtensionInfos:    file_goutput_proto_extTypes,
	}.Build()
	File_goutput_proto = out.File
	file_goutput_proto_rawDesc = nil
	file_goutput_proto_goTypes = nil
	file_goutput_proto_depIdxs = nil
}
