// SPDX-FileCopyrightText: Copyright (c) 2022 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: riva/proto/riva_common.proto

package riva_speech

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specifies the request ID of the request.
type RequestId struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestId) Reset() {
	*x = RequestId{}
	mi := &file_riva_proto_riva_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestId) ProtoMessage() {}

func (x *RequestId) ProtoReflect() protoreflect.Message {
	mi := &file_riva_proto_riva_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestId.ProtoReflect.Descriptor instead.
func (*RequestId) Descriptor() ([]byte, []int) {
	return file_riva_proto_riva_common_proto_rawDescGZIP(), []int{0}
}

func (x *RequestId) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_riva_proto_riva_common_proto protoreflect.FileDescriptor

var file_riva_proto_riva_common_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x72, 0x69, 0x76, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x69, 0x76,
	0x61, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x72, 0x69, 0x76, 0x61, 0x22, 0x21, 0x0a, 0x09, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1b,
	0x5a, 0x16, 0x6e, 0x76, 0x69, 0x64, 0x69, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x76,
	0x61, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_riva_proto_riva_common_proto_rawDescOnce sync.Once
	file_riva_proto_riva_common_proto_rawDescData []byte
)

func file_riva_proto_riva_common_proto_rawDescGZIP() []byte {
	file_riva_proto_riva_common_proto_rawDescOnce.Do(func() {
		file_riva_proto_riva_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_riva_proto_riva_common_proto_rawDesc), len(file_riva_proto_riva_common_proto_rawDesc)))
	})
	return file_riva_proto_riva_common_proto_rawDescData
}

var file_riva_proto_riva_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_riva_proto_riva_common_proto_goTypes = []any{
	(*RequestId)(nil), // 0: nvidia.riva.RequestId
}
var file_riva_proto_riva_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_riva_proto_riva_common_proto_init() }
func file_riva_proto_riva_common_proto_init() {
	if File_riva_proto_riva_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_riva_proto_riva_common_proto_rawDesc), len(file_riva_proto_riva_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_riva_proto_riva_common_proto_goTypes,
		DependencyIndexes: file_riva_proto_riva_common_proto_depIdxs,
		MessageInfos:      file_riva_proto_riva_common_proto_msgTypes,
	}.Build()
	File_riva_proto_riva_common_proto = out.File
	file_riva_proto_riva_common_proto_goTypes = nil
	file_riva_proto_riva_common_proto_depIdxs = nil
}
