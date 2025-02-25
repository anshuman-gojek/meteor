// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: odpf/assets/v1beta2/group.proto

package assetsv1beta2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Group represents a group of users and resources.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The email of the group.
	// Example: `xyz@xyz.com`
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The members of the group.
	// For example look at schema of the member.
	Members []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// List of attributes the model has.
	Attributes *structpb.Struct `protobuf:"bytes,10,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_v1beta2_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_v1beta2_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_odpf_assets_v1beta2_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Group) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Group) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Member represents a user.
type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier for the user.
	// Example: `user:example`.
	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	// The role of the user.
	// Example: `owner`.
	Role string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_v1beta2_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_v1beta2_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_odpf_assets_v1beta2_group_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Member) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_odpf_assets_v1beta2_group_proto protoreflect.FileDescriptor

var file_odpf_assets_v1beta2_group_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x42, 0x51, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x42, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x64, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_assets_v1beta2_group_proto_rawDescOnce sync.Once
	file_odpf_assets_v1beta2_group_proto_rawDescData = file_odpf_assets_v1beta2_group_proto_rawDesc
)

func file_odpf_assets_v1beta2_group_proto_rawDescGZIP() []byte {
	file_odpf_assets_v1beta2_group_proto_rawDescOnce.Do(func() {
		file_odpf_assets_v1beta2_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_assets_v1beta2_group_proto_rawDescData)
	})
	return file_odpf_assets_v1beta2_group_proto_rawDescData
}

var file_odpf_assets_v1beta2_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_odpf_assets_v1beta2_group_proto_goTypes = []interface{}{
	(*Group)(nil),           // 0: odpf.assets.v1beta2.Group
	(*Member)(nil),          // 1: odpf.assets.v1beta2.Member
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_odpf_assets_v1beta2_group_proto_depIdxs = []int32{
	1, // 0: odpf.assets.v1beta2.Group.members:type_name -> odpf.assets.v1beta2.Member
	2, // 1: odpf.assets.v1beta2.Group.attributes:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_odpf_assets_v1beta2_group_proto_init() }
func file_odpf_assets_v1beta2_group_proto_init() {
	if File_odpf_assets_v1beta2_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_assets_v1beta2_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_odpf_assets_v1beta2_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_odpf_assets_v1beta2_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_assets_v1beta2_group_proto_goTypes,
		DependencyIndexes: file_odpf_assets_v1beta2_group_proto_depIdxs,
		MessageInfos:      file_odpf_assets_v1beta2_group_proto_msgTypes,
	}.Build()
	File_odpf_assets_v1beta2_group_proto = out.File
	file_odpf_assets_v1beta2_group_proto_rawDesc = nil
	file_odpf_assets_v1beta2_group_proto_goTypes = nil
	file_odpf_assets_v1beta2_group_proto_depIdxs = nil
}
