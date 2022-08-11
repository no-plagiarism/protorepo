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
	UserUid    string `protobuf:"bytes,4,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
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

func (x *CheckText) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

type CheckTextResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckResult *CheckTextResult_CheckResult `protobuf:"bytes,3,opt,name=check_result,json=checkResult,proto3" json:"check_result,omitempty"`
	UserUid     string                       `protobuf:"bytes,4,opt,name=user_uid,json=userUid,proto3" json:"user_uid,omitempty"`
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

func (x *CheckTextResult) GetCheckResult() *CheckTextResult_CheckResult {
	if x != nil {
		return x.CheckResult
	}
	return nil
}

func (x *CheckTextResult) GetUserUid() string {
	if x != nil {
		return x.UserUid
	}
	return ""
}

type CheckTextResult_CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash             string                            `protobuf:"bytes,1,opt,name=file_hash,json=fileHash,proto3" json:"file_hash,omitempty"`
	FileName             string                            `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	PlagiarismPercentage float64                           `protobuf:"fixed64,3,opt,name=plagiarism_percentage,json=plagiarismPercentage,proto3" json:"plagiarism_percentage,omitempty"`
	PlagiarismResult     *CheckTextResult_PlagiarismResult `protobuf:"bytes,5,opt,name=plagiarism_result,json=plagiarismResult,proto3" json:"plagiarism_result,omitempty"`
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

func (x *CheckTextResult_CheckResult) GetPlagiarismResult() *CheckTextResult_PlagiarismResult {
	if x != nil {
		return x.PlagiarismResult
	}
	return nil
}

type CheckTextResult_PlagiarismResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimilarFiles       []*CheckTextResult_SimilarFile `protobuf:"bytes,1,rep,name=similar_files,json=similarFiles,proto3" json:"similar_files,omitempty"`
	Sequences          []*CheckTextResult_Sequence    `protobuf:"bytes,2,rep,name=sequences,proto3" json:"sequences,omitempty"`
	Urls               []string                       `protobuf:"bytes,3,rep,name=urls,proto3" json:"urls,omitempty"`
	Text               string                         `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	SymbolReplacements []int64                        `protobuf:"varint,5,rep,packed,name=symbol_replacements,json=symbolReplacements,proto3" json:"symbol_replacements,omitempty"`
}

func (x *CheckTextResult_PlagiarismResult) Reset() {
	*x = CheckTextResult_PlagiarismResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResult_PlagiarismResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResult_PlagiarismResult) ProtoMessage() {}

func (x *CheckTextResult_PlagiarismResult) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResult_PlagiarismResult.ProtoReflect.Descriptor instead.
func (*CheckTextResult_PlagiarismResult) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2, 1}
}

func (x *CheckTextResult_PlagiarismResult) GetSimilarFiles() []*CheckTextResult_SimilarFile {
	if x != nil {
		return x.SimilarFiles
	}
	return nil
}

func (x *CheckTextResult_PlagiarismResult) GetSequences() []*CheckTextResult_Sequence {
	if x != nil {
		return x.Sequences
	}
	return nil
}

func (x *CheckTextResult_PlagiarismResult) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *CheckTextResult_PlagiarismResult) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CheckTextResult_PlagiarismResult) GetSymbolReplacements() []int64 {
	if x != nil {
		return x.SymbolReplacements
	}
	return nil
}

type CheckTextResult_SimilarFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string  `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Url      *string `protobuf:"bytes,2,opt,name=url,proto3,oneof" json:"url,omitempty"`
}

func (x *CheckTextResult_SimilarFile) Reset() {
	*x = CheckTextResult_SimilarFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResult_SimilarFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResult_SimilarFile) ProtoMessage() {}

func (x *CheckTextResult_SimilarFile) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResult_SimilarFile.ProtoReflect.Descriptor instead.
func (*CheckTextResult_SimilarFile) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2, 2}
}

func (x *CheckTextResult_SimilarFile) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CheckTextResult_SimilarFile) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

type CheckTextResult_Sequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *CheckTextResult_Sequence) Reset() {
	*x = CheckTextResult_Sequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResult_Sequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResult_Sequence) ProtoMessage() {}

func (x *CheckTextResult_Sequence) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResult_Sequence.ProtoReflect.Descriptor instead.
func (*CheckTextResult_Sequence) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2, 3}
}

func (x *CheckTextResult_Sequence) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *CheckTextResult_Sequence) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
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
	0x6c, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x64,
	0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x55, 0x69, 0x64, 0x22, 0xdd, 0x05, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x55,
	0x69, 0x64, 0x1a, 0xda, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15,
	0x70, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x14, 0x70, 0x6c, 0x61,
	0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x5c, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x50, 0x6c, 0x61,
	0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x10, 0x70,
	0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a,
	0x83, 0x02, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x67, 0x69, 0x61, 0x72, 0x69, 0x73, 0x6d, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f, 0x0a, 0x0d, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x69, 0x6d, 0x69,
	0x6c, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0c, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x09, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x12, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x49, 0x0a, 0x0b, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x75, 0x72, 0x6c,
	0x1a, 0x32, 0x0a, 0x08, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_messages_proto_goTypes = []interface{}{
	(*AddFileRequest)(nil),                   // 0: kafka.message.AddFileRequest
	(*CheckText)(nil),                        // 1: kafka.message.CheckText
	(*CheckTextResult)(nil),                  // 2: kafka.message.CheckTextResult
	(*CheckTextResult_CheckResult)(nil),      // 3: kafka.message.CheckTextResult.CheckResult
	(*CheckTextResult_PlagiarismResult)(nil), // 4: kafka.message.CheckTextResult.PlagiarismResult
	(*CheckTextResult_SimilarFile)(nil),      // 5: kafka.message.CheckTextResult.SimilarFile
	(*CheckTextResult_Sequence)(nil),         // 6: kafka.message.CheckTextResult.Sequence
}
var file_messages_proto_depIdxs = []int32{
	3, // 0: kafka.message.CheckTextResult.check_result:type_name -> kafka.message.CheckTextResult.CheckResult
	4, // 1: kafka.message.CheckTextResult.CheckResult.plagiarism_result:type_name -> kafka.message.CheckTextResult.PlagiarismResult
	5, // 2: kafka.message.CheckTextResult.PlagiarismResult.similar_files:type_name -> kafka.message.CheckTextResult.SimilarFile
	6, // 3: kafka.message.CheckTextResult.PlagiarismResult.sequences:type_name -> kafka.message.CheckTextResult.Sequence
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResult_PlagiarismResult); i {
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
		file_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResult_SimilarFile); i {
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
		file_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResult_Sequence); i {
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
	file_messages_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
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
