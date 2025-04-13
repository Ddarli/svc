// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: service.proto

package fileservice

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

type UploadUserFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	MimeType      string                 `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	FileData      []byte                 `protobuf:"bytes,4,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadUserFileRequest) Reset() {
	*x = UploadUserFileRequest{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadUserFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUserFileRequest) ProtoMessage() {}

func (x *UploadUserFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUserFileRequest.ProtoReflect.Descriptor instead.
func (*UploadUserFileRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadUserFileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UploadUserFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadUserFileRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *UploadUserFileRequest) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *UploadUserFileRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UploadUserFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadUserFileResponse) Reset() {
	*x = UploadUserFileResponse{}
	mi := &file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadUserFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadUserFileResponse) ProtoMessage() {}

func (x *UploadUserFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadUserFileResponse.ProtoReflect.Descriptor instead.
func (*UploadUserFileResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadUserFileResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListUserFilesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserFilesRequest) Reset() {
	*x = ListUserFilesRequest{}
	mi := &file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserFilesRequest) ProtoMessage() {}

func (x *ListUserFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserFilesRequest.ProtoReflect.Descriptor instead.
func (*ListUserFilesRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserFilesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FileMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	mi := &file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *FileMetadata) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *FileMetadata) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FileMetadata) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ListUserFilesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*FileMetadata        `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserFilesResponse) Reset() {
	*x = ListUserFilesResponse{}
	mi := &file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserFilesResponse) ProtoMessage() {}

func (x *ListUserFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserFilesResponse.ProtoReflect.Descriptor instead.
func (*ListUserFilesResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListUserFilesResponse) GetFiles() []*FileMetadata {
	if x != nil {
		return x.Files
	}
	return nil
}

type DownloadUserFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FileId        string                 `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadUserFileRequest) Reset() {
	*x = DownloadUserFileRequest{}
	mi := &file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadUserFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadUserFileRequest) ProtoMessage() {}

func (x *DownloadUserFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadUserFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadUserFileRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadUserFileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DownloadUserFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DownloadUserFileResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileName      string                 `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	MimeType      string                 `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	FileData      []byte                 `protobuf:"bytes,3,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadUserFileResponse) Reset() {
	*x = DownloadUserFileResponse{}
	mi := &file_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadUserFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadUserFileResponse) ProtoMessage() {}

func (x *DownloadUserFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadUserFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadUserFileResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadUserFileResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *DownloadUserFileResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *DownloadUserFileResponse) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

const file_service_proto_rawDesc = "" +
	"\n" +
	"\rservice.proto\x12\vfileservice\"\xa9\x01\n" +
	"\x15UploadUserFileRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1b\n" +
	"\tfile_name\x18\x02 \x01(\tR\bfileName\x12\x1b\n" +
	"\tmime_type\x18\x03 \x01(\tR\bmimeType\x12\x1b\n" +
	"\tfile_data\x18\x04 \x01(\fR\bfileData\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\"2\n" +
	"\x16UploadUserFileResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"/\n" +
	"\x14ListUserFilesRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"~\n" +
	"\fFileMetadata\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x1b\n" +
	"\tfile_name\x18\x02 \x01(\tR\bfileName\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\"H\n" +
	"\x15ListUserFilesResponse\x12/\n" +
	"\x05files\x18\x01 \x03(\v2\x19.fileservice.FileMetadataR\x05files\"K\n" +
	"\x17DownloadUserFileRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x17\n" +
	"\afile_id\x18\x02 \x01(\tR\x06fileId\"q\n" +
	"\x18DownloadUserFileResponse\x12\x1b\n" +
	"\tfile_name\x18\x01 \x01(\tR\bfileName\x12\x1b\n" +
	"\tmime_type\x18\x02 \x01(\tR\bmimeType\x12\x1b\n" +
	"\tfile_data\x18\x03 \x01(\fR\bfileData2\xa1\x02\n" +
	"\vFileService\x12Y\n" +
	"\x0eUploadUserFile\x12\".fileservice.UploadUserFileRequest\x1a#.fileservice.UploadUserFileResponse\x12V\n" +
	"\rListUserFiles\x12!.fileservice.ListUserFilesRequest\x1a\".fileservice.ListUserFilesResponse\x12_\n" +
	"\x10DownloadUserFile\x12$.fileservice.DownloadUserFileRequest\x1a%.fileservice.DownloadUserFileResponseB\"Z data-processor/proto/fileserviceb\x06proto3"

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData []byte
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)))
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_service_proto_goTypes = []any{
	(*UploadUserFileRequest)(nil),    // 0: fileservice.UploadUserFileRequest
	(*UploadUserFileResponse)(nil),   // 1: fileservice.UploadUserFileResponse
	(*ListUserFilesRequest)(nil),     // 2: fileservice.ListUserFilesRequest
	(*FileMetadata)(nil),             // 3: fileservice.FileMetadata
	(*ListUserFilesResponse)(nil),    // 4: fileservice.ListUserFilesResponse
	(*DownloadUserFileRequest)(nil),  // 5: fileservice.DownloadUserFileRequest
	(*DownloadUserFileResponse)(nil), // 6: fileservice.DownloadUserFileResponse
}
var file_service_proto_depIdxs = []int32{
	3, // 0: fileservice.ListUserFilesResponse.files:type_name -> fileservice.FileMetadata
	0, // 1: fileservice.FileService.UploadUserFile:input_type -> fileservice.UploadUserFileRequest
	2, // 2: fileservice.FileService.ListUserFiles:input_type -> fileservice.ListUserFilesRequest
	5, // 3: fileservice.FileService.DownloadUserFile:input_type -> fileservice.DownloadUserFileRequest
	1, // 4: fileservice.FileService.UploadUserFile:output_type -> fileservice.UploadUserFileResponse
	4, // 5: fileservice.FileService.ListUserFiles:output_type -> fileservice.ListUserFilesResponse
	6, // 6: fileservice.FileService.DownloadUserFile:output_type -> fileservice.DownloadUserFileResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
