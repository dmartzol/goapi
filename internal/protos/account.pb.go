// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: internal/protos/account.proto

package templateprotos

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

type AccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccountRequest) Reset() {
	*x = AccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountRequest) ProtoMessage() {}

func (x *AccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountRequest.ProtoReflect.Descriptor instead.
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return file_internal_protos_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateAccountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string                 `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string                 `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Dob       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender    string                 `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Password  string                 `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Email     string                 `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateAccountMessage) Reset() {
	*x = CreateAccountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountMessage) ProtoMessage() {}

func (x *CreateAccountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountMessage.ProtoReflect.Descriptor instead.
func (*CreateAccountMessage) Descriptor() ([]byte, []int) {
	return file_internal_protos_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateAccountMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateAccountMessage) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *CreateAccountMessage) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateAccountMessage) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateAccountMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AccountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	FirstName   string                 `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string                 `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Dob         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender      string                 `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	PassHash    string                 `protobuf:"bytes,8,opt,name=pass_hash,json=passHash,proto3" json:"pass_hash,omitempty"`
	Email       string                 `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AccountMessage) Reset() {
	*x = AccountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protos_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMessage) ProtoMessage() {}

func (x *AccountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protos_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountMessage.ProtoReflect.Descriptor instead.
func (*AccountMessage) Descriptor() ([]byte, []int) {
	return file_internal_protos_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountMessage) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *AccountMessage) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *AccountMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AccountMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AccountMessage) GetDob() *timestamppb.Timestamp {
	if x != nil {
		return x.Dob
	}
	return nil
}

func (x *AccountMessage) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *AccountMessage) GetPassHash() string {
	if x != nil {
		return x.PassHash
	}
	return ""
}

func (x *AccountMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_internal_protos_account_proto protoreflect.FileDescriptor

var file_internal_protos_account_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x20, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0xd3, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x03,
	0x64, 0x6f, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x32, 0xb0, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x57, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x07, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6d, 0x61, 0x72, 0x74, 0x7a, 0x6f, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protos_account_proto_rawDescOnce sync.Once
	file_internal_protos_account_proto_rawDescData = file_internal_protos_account_proto_rawDesc
)

func file_internal_protos_account_proto_rawDescGZIP() []byte {
	file_internal_protos_account_proto_rawDescOnce.Do(func() {
		file_internal_protos_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protos_account_proto_rawDescData)
	})
	return file_internal_protos_account_proto_rawDescData
}

var file_internal_protos_account_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_protos_account_proto_goTypes = []interface{}{
	(*AccountRequest)(nil),        // 0: templateprotos.AccountRequest
	(*CreateAccountMessage)(nil),  // 1: templateprotos.CreateAccountMessage
	(*AccountMessage)(nil),        // 2: templateprotos.AccountMessage
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_internal_protos_account_proto_depIdxs = []int32{
	3, // 0: templateprotos.CreateAccountMessage.dob:type_name -> google.protobuf.Timestamp
	3, // 1: templateprotos.AccountMessage.created_time:type_name -> google.protobuf.Timestamp
	3, // 2: templateprotos.AccountMessage.updated_time:type_name -> google.protobuf.Timestamp
	3, // 3: templateprotos.AccountMessage.dob:type_name -> google.protobuf.Timestamp
	1, // 4: templateprotos.Accounts.CreateAccount:input_type -> templateprotos.CreateAccountMessage
	0, // 5: templateprotos.Accounts.Account:input_type -> templateprotos.AccountRequest
	2, // 6: templateprotos.Accounts.CreateAccount:output_type -> templateprotos.AccountMessage
	2, // 7: templateprotos.Accounts.Account:output_type -> templateprotos.AccountMessage
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_protos_account_proto_init() }
func file_internal_protos_account_proto_init() {
	if File_internal_protos_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protos_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountRequest); i {
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
		file_internal_protos_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountMessage); i {
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
		file_internal_protos_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountMessage); i {
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
			RawDescriptor: file_internal_protos_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protos_account_proto_goTypes,
		DependencyIndexes: file_internal_protos_account_proto_depIdxs,
		MessageInfos:      file_internal_protos_account_proto_msgTypes,
	}.Build()
	File_internal_protos_account_proto = out.File
	file_internal_protos_account_proto_rawDesc = nil
	file_internal_protos_account_proto_goTypes = nil
	file_internal_protos_account_proto_depIdxs = nil
}
