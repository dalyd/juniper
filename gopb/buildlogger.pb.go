// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: buildlogger.proto

package gopb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogStorage int32

const (
	LogStorage_LOG_STORAGE_S3     LogStorage = 0
	LogStorage_LOG_STORAGE_GRIDFS LogStorage = 1
	LogStorage_LOG_STORAGE_LOCAL  LogStorage = 2
)

// Enum value maps for LogStorage.
var (
	LogStorage_name = map[int32]string{
		0: "LOG_STORAGE_S3",
		1: "LOG_STORAGE_GRIDFS",
		2: "LOG_STORAGE_LOCAL",
	}
	LogStorage_value = map[string]int32{
		"LOG_STORAGE_S3":     0,
		"LOG_STORAGE_GRIDFS": 1,
		"LOG_STORAGE_LOCAL":  2,
	}
)

func (x LogStorage) Enum() *LogStorage {
	p := new(LogStorage)
	*p = x
	return p
}

func (x LogStorage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogStorage) Descriptor() protoreflect.EnumDescriptor {
	return file_buildlogger_proto_enumTypes[0].Descriptor()
}

func (LogStorage) Type() protoreflect.EnumType {
	return &file_buildlogger_proto_enumTypes[0]
}

func (x LogStorage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogStorage.Descriptor instead.
func (LogStorage) EnumDescriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{0}
}

type LogFormat int32

const (
	LogFormat_LOG_FORMAT_UNKNOWN LogFormat = 0
	LogFormat_LOG_FORMAT_TEXT    LogFormat = 1
	LogFormat_LOG_FORMAT_JSON    LogFormat = 2
	LogFormat_LOG_FORMAT_BSON    LogFormat = 3
)

// Enum value maps for LogFormat.
var (
	LogFormat_name = map[int32]string{
		0: "LOG_FORMAT_UNKNOWN",
		1: "LOG_FORMAT_TEXT",
		2: "LOG_FORMAT_JSON",
		3: "LOG_FORMAT_BSON",
	}
	LogFormat_value = map[string]int32{
		"LOG_FORMAT_UNKNOWN": 0,
		"LOG_FORMAT_TEXT":    1,
		"LOG_FORMAT_JSON":    2,
		"LOG_FORMAT_BSON":    3,
	}
)

func (x LogFormat) Enum() *LogFormat {
	p := new(LogFormat)
	*p = x
	return p
}

func (x LogFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_buildlogger_proto_enumTypes[1].Descriptor()
}

func (LogFormat) Type() protoreflect.EnumType {
	return &file_buildlogger_proto_enumTypes[1]
}

func (x LogFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogFormat.Descriptor instead.
func (LogFormat) EnumDescriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{1}
}

type LogData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *LogInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Storage LogStorage `protobuf:"varint,2,opt,name=storage,proto3,enum=cedar.LogStorage" json:"storage,omitempty"`
}

func (x *LogData) Reset() {
	*x = LogData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogData) ProtoMessage() {}

func (x *LogData) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogData.ProtoReflect.Descriptor instead.
func (*LogData) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{0}
}

func (x *LogData) GetInfo() *LogInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *LogData) GetStorage() LogStorage {
	if x != nil {
		return x.Storage
	}
	return LogStorage_LOG_STORAGE_S3
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project   string            `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variant   string            `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	TaskName  string            `protobuf:"bytes,4,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskId    string            `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Execution int32             `protobuf:"varint,6,opt,name=execution,proto3" json:"execution,omitempty"`
	TestName  string            `protobuf:"bytes,7,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	Trial     int32             `protobuf:"varint,8,opt,name=trial,proto3" json:"trial,omitempty"`
	ProcName  string            `protobuf:"bytes,9,opt,name=proc_name,json=procName,proto3" json:"proc_name,omitempty"`
	Format    LogFormat         `protobuf:"varint,10,opt,name=format,proto3,enum=cedar.LogFormat" json:"format,omitempty"`
	Tags      []string          `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty"`
	Arguments map[string]string `protobuf:"bytes,12,rep,name=arguments,proto3" json:"arguments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mainline  bool              `protobuf:"varint,13,opt,name=mainline,proto3" json:"mainline,omitempty"`
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{1}
}

func (x *LogInfo) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *LogInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LogInfo) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *LogInfo) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *LogInfo) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *LogInfo) GetExecution() int32 {
	if x != nil {
		return x.Execution
	}
	return 0
}

func (x *LogInfo) GetTestName() string {
	if x != nil {
		return x.TestName
	}
	return ""
}

func (x *LogInfo) GetTrial() int32 {
	if x != nil {
		return x.Trial
	}
	return 0
}

func (x *LogInfo) GetProcName() string {
	if x != nil {
		return x.ProcName
	}
	return ""
}

func (x *LogInfo) GetFormat() LogFormat {
	if x != nil {
		return x.Format
	}
	return LogFormat_LOG_FORMAT_UNKNOWN
}

func (x *LogInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *LogInfo) GetArguments() map[string]string {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *LogInfo) GetMainline() bool {
	if x != nil {
		return x.Mainline
	}
	return false
}

type LogLines struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId string     `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Lines []*LogLine `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *LogLines) Reset() {
	*x = LogLines{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLines) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLines) ProtoMessage() {}

func (x *LogLines) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLines.ProtoReflect.Descriptor instead.
func (*LogLines) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{2}
}

func (x *LogLines) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *LogLines) GetLines() []*LogLine {
	if x != nil {
		return x.Lines
	}
	return nil
}

type LogLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority  int32                  `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LogLine) Reset() {
	*x = LogLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLine) ProtoMessage() {}

func (x *LogLine) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLine.ProtoReflect.Descriptor instead.
func (*LogLine) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{3}
}

func (x *LogLine) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *LogLine) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *LogLine) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type LogEndInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId    string `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	ExitCode int32  `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *LogEndInfo) Reset() {
	*x = LogEndInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEndInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEndInfo) ProtoMessage() {}

func (x *LogEndInfo) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEndInfo.ProtoReflect.Descriptor instead.
func (*LogEndInfo) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{4}
}

func (x *LogEndInfo) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

func (x *LogEndInfo) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

type BuildloggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId string `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
}

func (x *BuildloggerResponse) Reset() {
	*x = BuildloggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildlogger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildloggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildloggerResponse) ProtoMessage() {}

func (x *BuildloggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildlogger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildloggerResponse.ProtoReflect.Descriptor instead.
func (*BuildloggerResponse) Descriptor() ([]byte, []int) {
	return file_buildlogger_proto_rawDescGZIP(), []int{5}
}

func (x *BuildloggerResponse) GetLogId() string {
	if x != nil {
		return x.LogId
	}
	return ""
}

var File_buildlogger_proto protoreflect.FileDescriptor

var file_buildlogger_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x65, 0x64, 0x61, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x07, 0x4c,
	0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x65,
	0x64, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xd0, 0x03, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x3b, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x3c, 0x0a, 0x0e,
	0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47, 0x0a, 0x08, 0x4c, 0x6f,
	0x67, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x65, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x22, 0x73, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x45,
	0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x13, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x2a, 0x4f, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x33, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f,
	0x47, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47, 0x45, 0x5f, 0x47, 0x52, 0x49, 0x44, 0x46, 0x53,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x41, 0x47,
	0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x02, 0x2a, 0x62, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x46, 0x4f,
	0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x54, 0x45, 0x58,
	0x54, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x42, 0x53, 0x4f, 0x4e, 0x10, 0x03, 0x32, 0x81, 0x02,
	0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x37, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x2e, 0x63, 0x65, 0x64,
	0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x64,
	0x61, 0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72,
	0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x64, 0x61,
	0x72, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x39, 0x0a, 0x08, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4c,
	0x6f, 0x67, 0x12, 0x11, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x63, 0x65, 0x64, 0x61, 0x72, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_buildlogger_proto_rawDescOnce sync.Once
	file_buildlogger_proto_rawDescData = file_buildlogger_proto_rawDesc
)

func file_buildlogger_proto_rawDescGZIP() []byte {
	file_buildlogger_proto_rawDescOnce.Do(func() {
		file_buildlogger_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildlogger_proto_rawDescData)
	})
	return file_buildlogger_proto_rawDescData
}

var file_buildlogger_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_buildlogger_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_buildlogger_proto_goTypes = []interface{}{
	(LogStorage)(0),               // 0: cedar.LogStorage
	(LogFormat)(0),                // 1: cedar.LogFormat
	(*LogData)(nil),               // 2: cedar.LogData
	(*LogInfo)(nil),               // 3: cedar.LogInfo
	(*LogLines)(nil),              // 4: cedar.LogLines
	(*LogLine)(nil),               // 5: cedar.LogLine
	(*LogEndInfo)(nil),            // 6: cedar.LogEndInfo
	(*BuildloggerResponse)(nil),   // 7: cedar.BuildloggerResponse
	nil,                           // 8: cedar.LogInfo.ArgumentsEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_buildlogger_proto_depIdxs = []int32{
	3,  // 0: cedar.LogData.info:type_name -> cedar.LogInfo
	0,  // 1: cedar.LogData.storage:type_name -> cedar.LogStorage
	1,  // 2: cedar.LogInfo.format:type_name -> cedar.LogFormat
	8,  // 3: cedar.LogInfo.arguments:type_name -> cedar.LogInfo.ArgumentsEntry
	5,  // 4: cedar.LogLines.lines:type_name -> cedar.LogLine
	9,  // 5: cedar.LogLine.timestamp:type_name -> google.protobuf.Timestamp
	2,  // 6: cedar.Buildlogger.CreateLog:input_type -> cedar.LogData
	4,  // 7: cedar.Buildlogger.AppendLogLines:input_type -> cedar.LogLines
	4,  // 8: cedar.Buildlogger.StreamLogLines:input_type -> cedar.LogLines
	6,  // 9: cedar.Buildlogger.CloseLog:input_type -> cedar.LogEndInfo
	7,  // 10: cedar.Buildlogger.CreateLog:output_type -> cedar.BuildloggerResponse
	7,  // 11: cedar.Buildlogger.AppendLogLines:output_type -> cedar.BuildloggerResponse
	7,  // 12: cedar.Buildlogger.StreamLogLines:output_type -> cedar.BuildloggerResponse
	7,  // 13: cedar.Buildlogger.CloseLog:output_type -> cedar.BuildloggerResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_buildlogger_proto_init() }
func file_buildlogger_proto_init() {
	if File_buildlogger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buildlogger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogData); i {
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
		file_buildlogger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
		file_buildlogger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLines); i {
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
		file_buildlogger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLine); i {
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
		file_buildlogger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEndInfo); i {
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
		file_buildlogger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildloggerResponse); i {
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
			RawDescriptor: file_buildlogger_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildlogger_proto_goTypes,
		DependencyIndexes: file_buildlogger_proto_depIdxs,
		EnumInfos:         file_buildlogger_proto_enumTypes,
		MessageInfos:      file_buildlogger_proto_msgTypes,
	}.Build()
	File_buildlogger_proto = out.File
	file_buildlogger_proto_rawDesc = nil
	file_buildlogger_proto_goTypes = nil
	file_buildlogger_proto_depIdxs = nil
}
