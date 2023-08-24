// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: test.proto

package pb

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

// 做测试
type ChoiceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ChoiceItem) Reset() {
	*x = ChoiceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChoiceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChoiceItem) ProtoMessage() {}

func (x *ChoiceItem) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChoiceItem.ProtoReflect.Descriptor instead.
func (*ChoiceItem) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *ChoiceItem) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ChoiceItem) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type TestReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Choice []*ChoiceItem `protobuf:"bytes,1,rep,name=Choice,proto3" json:"Choice,omitempty"`
	Token  string        `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *TestReq) Reset() {
	*x = TestReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestReq) ProtoMessage() {}

func (x *TestReq) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestReq.ProtoReflect.Descriptor instead.
func (*TestReq) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestReq) GetChoice() []*ChoiceItem {
	if x != nil {
		return x.Choice
	}
	return nil
}

func (x *TestReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TestResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag   bool   `protobuf:"varint,1,opt,name=Flag,proto3" json:"Flag,omitempty"`
	Userid string `protobuf:"bytes,2,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *TestResp) Reset() {
	*x = TestResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResp) ProtoMessage() {}

func (x *TestResp) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResp.ProtoReflect.Descriptor instead.
func (*TestResp) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestResp) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

func (x *TestResp) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

// 返回测试结果
type TestInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *TestInfoReq) Reset() {
	*x = TestInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestInfoReq) ProtoMessage() {}

func (x *TestInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestInfoReq.ProtoReflect.Descriptor instead.
func (*TestInfoReq) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestInfoReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TestInfoReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type TestInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Gender      string   `protobuf:"bytes,2,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Major       string   `protobuf:"bytes,3,opt,name=Major,proto3" json:"Major,omitempty"`
	Grade       string   `protobuf:"bytes,4,opt,name=Grade,proto3" json:"Grade,omitempty"`
	LeQunXing   int64    `protobuf:"varint,5,opt,name=LeQunXing,proto3" json:"LeQunXing,omitempty"`
	YouHengXing int64    `protobuf:"varint,6,opt,name=YouHengXing,proto3" json:"YouHengXing,omitempty"`
	XingFenXing int64    `protobuf:"varint,7,opt,name=XingFenXing,proto3" json:"XingFenXing,omitempty"`
	CongHuiXing int64    `protobuf:"varint,8,opt,name=CongHuiXing,proto3" json:"CongHuiXing,omitempty"`
	JiaoJiXing  int64    `protobuf:"varint,9,opt,name=JiaoJiXing,proto3" json:"JiaoJiXing,omitempty"`
	HuaiYiXing  int64    `protobuf:"varint,10,opt,name=HuaiYiXing,proto3" json:"HuaiYiXing,omitempty"`
	WenDingXing int64    `protobuf:"varint,11,opt,name=WenDingXing,proto3" json:"WenDingXing,omitempty"`
	Choice      []string `protobuf:"bytes,12,rep,name=Choice,proto3" json:"Choice,omitempty"`
}

func (x *TestInfoResp) Reset() {
	*x = TestInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestInfoResp) ProtoMessage() {}

func (x *TestInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestInfoResp.ProtoReflect.Descriptor instead.
func (*TestInfoResp) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *TestInfoResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestInfoResp) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *TestInfoResp) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *TestInfoResp) GetGrade() string {
	if x != nil {
		return x.Grade
	}
	return ""
}

func (x *TestInfoResp) GetLeQunXing() int64 {
	if x != nil {
		return x.LeQunXing
	}
	return 0
}

func (x *TestInfoResp) GetYouHengXing() int64 {
	if x != nil {
		return x.YouHengXing
	}
	return 0
}

func (x *TestInfoResp) GetXingFenXing() int64 {
	if x != nil {
		return x.XingFenXing
	}
	return 0
}

func (x *TestInfoResp) GetCongHuiXing() int64 {
	if x != nil {
		return x.CongHuiXing
	}
	return 0
}

func (x *TestInfoResp) GetJiaoJiXing() int64 {
	if x != nil {
		return x.JiaoJiXing
	}
	return 0
}

func (x *TestInfoResp) GetHuaiYiXing() int64 {
	if x != nil {
		return x.HuaiYiXing
	}
	return 0
}

func (x *TestInfoResp) GetWenDingXing() int64 {
	if x != nil {
		return x.WenDingXing
	}
	return 0
}

func (x *TestInfoResp) GetChoice() []string {
	if x != nil {
		return x.Choice
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x22, 0x38, 0x0a, 0x0a, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x07,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x43, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22,
	0x3b, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xe4, 0x02, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x61, 0x6a,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x61, 0x6a, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x65, 0x51, 0x75, 0x6e, 0x58, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4c, 0x65, 0x51, 0x75, 0x6e, 0x58,
	0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x59, 0x6f, 0x75, 0x48, 0x65, 0x6e, 0x67, 0x58, 0x69,
	0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x59, 0x6f, 0x75, 0x48, 0x65, 0x6e,
	0x67, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x58, 0x69, 0x6e, 0x67, 0x46, 0x65, 0x6e,
	0x58, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x58, 0x69, 0x6e, 0x67,
	0x46, 0x65, 0x6e, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x67, 0x48,
	0x75, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x43, 0x6f,
	0x6e, 0x67, 0x48, 0x75, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x4a, 0x69, 0x61,
	0x6f, 0x4a, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x4a,
	0x69, 0x61, 0x6f, 0x4a, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x48, 0x75, 0x61,
	0x69, 0x59, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x48,
	0x75, 0x61, 0x69, 0x59, 0x69, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x65, 0x6e,
	0x44, 0x69, 0x6e, 0x67, 0x58, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x57, 0x65, 0x6e, 0x44, 0x69, 0x6e, 0x67, 0x58, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x43, 0x68, 0x6f,
	0x69, 0x63, 0x65, 0x32, 0x71, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0d, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0f,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x11, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_test_proto_goTypes = []interface{}{
	(*ChoiceItem)(nil),   // 0: test.ChoiceItem
	(*TestReq)(nil),      // 1: test.TestReq
	(*TestResp)(nil),     // 2: test.TestResp
	(*TestInfoReq)(nil),  // 3: test.TestInfoReq
	(*TestInfoResp)(nil), // 4: test.TestInfoResp
}
var file_test_proto_depIdxs = []int32{
	0, // 0: test.TestReq.Choice:type_name -> test.ChoiceItem
	1, // 1: test.TestClient.UserTest:input_type -> test.TestReq
	3, // 2: test.TestClient.CheckTestResult:input_type -> test.TestInfoReq
	2, // 3: test.TestClient.UserTest:output_type -> test.TestResp
	4, // 4: test.TestClient.CheckTestResult:output_type -> test.TestInfoResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChoiceItem); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestReq); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResp); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestInfoReq); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestInfoResp); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
