// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: messages.proto

package messages

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

type AddFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// hash of the file
	Hash string  `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Url  *string `protobuf:"bytes,3,opt,name=url,proto3,oneof" json:"url,omitempty"`
	// full_path is the full path of the file on the local filesystem
	FullPath string `protobuf:"bytes,4,opt,name=full_path,json=fullPath,proto3" json:"full_path,omitempty"`
}

func (x *AddFileRequest) Reset() {
	*x = AddFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFileRequest) ProtoMessage() {}

func (x *AddFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFileRequest.ProtoReflect.Descriptor instead.
func (*AddFileRequest) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AddFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *AddFileRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *AddFileRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *AddFileRequest) GetFullPath() string {
	if x != nil {
		return x.FullPath
	}
	return ""
}

type CheckText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsGoogleOn bool   `protobuf:"varint,1,opt,name=is_google_on,json=isGoogleOn,proto3" json:"is_google_on,omitempty"`
	Document   string `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	UserId     int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckText) Reset() {
	*x = CheckText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckText) ProtoMessage() {}

func (x *CheckText) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckText.ProtoReflect.Descriptor instead.
func (*CheckText) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CheckText) GetIsGoogleOn() bool {
	if x != nil {
		return x.IsGoogleOn
	}
	return false
}

func (x *CheckText) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

func (x *CheckText) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CheckTextResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          int64                        `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CheckResult     *CheckTextResult_CheckResult `protobuf:"bytes,3,opt,name=check_result,json=checkResult,proto3" json:"check_result,omitempty"`
	CheckResultUuid string                       `protobuf:"bytes,4,opt,name=check_result_uuid,json=checkResultUuid,proto3" json:"check_result_uuid,omitempty"`
}

func (x *CheckTextResult) Reset() {
	*x = CheckTextResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResult) ProtoMessage() {}

func (x *CheckTextResult) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResult.ProtoReflect.Descriptor instead.
func (*CheckTextResult) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CheckTextResult) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckTextResult) GetCheckResult() *CheckTextResult_CheckResult {
	if x != nil {
		return x.CheckResult
	}
	return nil
}

func (x *CheckTextResult) GetCheckResultUuid() string {
	if x != nil {
		return x.CheckResultUuid
	}
	return ""
}

type CheckTextResult_CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash             string  `protobuf:"bytes,1,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	FileName             string  `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	PlagiarismPercentage float64 `protobuf:"fixed64,3,opt,name=plagiarism_percentage,json=plagiarismPercentage,proto3" json:"plagiarism_percentage,omitempty"`
	ExtendedData         string  `protobuf:"bytes,4,opt,name=extended_data,json=extendedData,proto3" json:"extended_data,omitempty"`
}

func (x *CheckTextResult_CheckResult) Reset() {
	*x = CheckTextResult_CheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResult_CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResult_CheckResult) ProtoMessage() {}

func (x *CheckTextResult_CheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResult_CheckResult.ProtoReflect.Descriptor instead.
func (*CheckTextResult_CheckResult) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CheckTextResult_CheckResult) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *CheckTextResult_CheckResult) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CheckTextResult_CheckResult) GetPlagiarismPercentage() float64 {
	if x != nil {
		return x.PlagiarismPercentage
	}
	return 0
}

func (x *CheckTextResult_CheckResult) GetExtendedData() string {
	if x != nil {
		return x.ExtendedData
	}
	return ""
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x7d, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c,
	0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x62,
	0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xc9, 0x02, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x4d, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55, 0x75, 0x69, 0x64, 0x1a, 0xa1, 0x01, 0x0a, 0x0b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x70, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69,
	0x73, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x14, 0x70, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_goTypes = []interface{}{
	(*AddFileRequest)(nil),              // 0: kafka.message.AddFileRequest
	(*CheckText)(nil),                   // 1: kafka.message.CheckText
	(*CheckTextResult)(nil),             // 2: kafka.message.CheckTextResult
	(*CheckTextResult_CheckResult)(nil), // 3: kafka.message.CheckTextResult.CheckResult
}
var file_messages_proto_depIdxs = []int32{
	3, // 0: kafka.message.CheckTextResult.check_result:type_name -> kafka.message.CheckTextResult.CheckResult
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFileRequest); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckText); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResult); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResult_CheckResult); i {
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
	file_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
