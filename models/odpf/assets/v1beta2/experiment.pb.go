// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: odpf/assets/v1beta2/experiment.proto

package assetsv1beta2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// An experiment is the set of configurations and filters that allow for
// systematically varying some independent variables to impact some other
// dependent variables.
type Experiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional: Type of the entity being experimented over. ex: customer, session,
	// device, driver etc.
	Entity string `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// Optional: Percentage of the traffic that the experiment is enabled for.
	TrafficPercent float32 `protobuf:"fixed32,2,opt,name=traffic_percent,json=trafficPercent,proto3" json:"traffic_percent,omitempty"`
	// The variants of the experiment possibly including the control group.
	Variants []*Experiment_Variant `protobuf:"bytes,3,rep,name=variants,proto3" json:"variants,omitempty"`
	// Optional: List of attributes the experiment has. This could include the
	// following:
	// - client_id[string]: The ID if the client running the experiment.
	// - client_name[string]: The name of the client running the experiment.
	// - primary_metric[string]: Used to determine a statistically significant
	//   winning or losing variant.
	// - guardrail_metric[string]: Business metric designed to indirectly measure
	//   business value and track any potentially misleading or erroneous results
	//   and analysis.
	// - variant_sample_size[double]: Sample size per variant.
	// - filter_rules[repeated string]: Textual representation of rules required
	//   to be satisfied for experiment to be shown to the user.
	// - start_time[RFC 3339 string]: The timestamp at which the
	//   experiment would start.
	// - end_time[RFC 3339 string]: The timestamp at which the
	//   experiment would end.
	Attributes *structpb.Struct `protobuf:"bytes,5,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The timestamp of the experiment's creation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The timestamp when the experiment was last modified.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,102,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Experiment) Reset() {
	*x = Experiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_v1beta2_experiment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experiment) ProtoMessage() {}

func (x *Experiment) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_v1beta2_experiment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experiment.ProtoReflect.Descriptor instead.
func (*Experiment) Descriptor() ([]byte, []int) {
	return file_odpf_assets_v1beta2_experiment_proto_rawDescGZIP(), []int{0}
}

func (x *Experiment) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *Experiment) GetTrafficPercent() float32 {
	if x != nil {
		return x.TrafficPercent
	}
	return 0
}

func (x *Experiment) GetVariants() []*Experiment_Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

func (x *Experiment) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Experiment) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Experiment) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Instance of configurations to be compared in the experiment.
type Experiment_Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the experiment variant.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Traffic percent enabled for the variant.
	TrafficPercent float32 `protobuf:"fixed32,2,opt,name=traffic_percent,json=trafficPercent,proto3" json:"traffic_percent,omitempty"`
	// Indicated whether the variant is the control for the experiment.
	IsControl bool `protobuf:"varint,3,opt,name=is_control,json=isControl,proto3" json:"is_control,omitempty"`
	// List of properties the entity has.
	Attributes *structpb.Struct `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Whether the variant has been promoted to all users.
	IsPromoted bool `protobuf:"varint,5,opt,name=is_promoted,json=isPromoted,proto3" json:"is_promoted,omitempty"`
}

func (x *Experiment_Variant) Reset() {
	*x = Experiment_Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_assets_v1beta2_experiment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Experiment_Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Experiment_Variant) ProtoMessage() {}

func (x *Experiment_Variant) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_assets_v1beta2_experiment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Experiment_Variant.ProtoReflect.Descriptor instead.
func (*Experiment_Variant) Descriptor() ([]byte, []int) {
	return file_odpf_assets_v1beta2_experiment_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Experiment_Variant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Experiment_Variant) GetTrafficPercent() float32 {
	if x != nil {
		return x.TrafficPercent
	}
	return 0
}

func (x *Experiment_Variant) GetIsControl() bool {
	if x != nil {
		return x.IsControl
	}
	return false
}

func (x *Experiment_Variant) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Experiment_Variant) GetIsPromoted() bool {
	if x != nil {
		return x.IsPromoted
	}
	return false
}

var File_odpf_assets_v1beta2_experiment_proto protoreflect.FileDescriptor

var file_odpf_assets_v1beta2_experiment_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x04, 0x0a, 0x0a, 0x45,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x08, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6f,
	0x64, 0x70, 0x66, 0x2e, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x32, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x12,
	0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x1a, 0xbf, 0x01, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x65, 0x64, 0x42, 0x56, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x42, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x3b, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_assets_v1beta2_experiment_proto_rawDescOnce sync.Once
	file_odpf_assets_v1beta2_experiment_proto_rawDescData = file_odpf_assets_v1beta2_experiment_proto_rawDesc
)

func file_odpf_assets_v1beta2_experiment_proto_rawDescGZIP() []byte {
	file_odpf_assets_v1beta2_experiment_proto_rawDescOnce.Do(func() {
		file_odpf_assets_v1beta2_experiment_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_assets_v1beta2_experiment_proto_rawDescData)
	})
	return file_odpf_assets_v1beta2_experiment_proto_rawDescData
}

var file_odpf_assets_v1beta2_experiment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_odpf_assets_v1beta2_experiment_proto_goTypes = []interface{}{
	(*Experiment)(nil),            // 0: odpf.assets.v1beta2.Experiment
	(*Experiment_Variant)(nil),    // 1: odpf.assets.v1beta2.Experiment.Variant
	(*structpb.Struct)(nil),       // 2: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_odpf_assets_v1beta2_experiment_proto_depIdxs = []int32{
	1, // 0: odpf.assets.v1beta2.Experiment.variants:type_name -> odpf.assets.v1beta2.Experiment.Variant
	2, // 1: odpf.assets.v1beta2.Experiment.attributes:type_name -> google.protobuf.Struct
	3, // 2: odpf.assets.v1beta2.Experiment.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: odpf.assets.v1beta2.Experiment.update_time:type_name -> google.protobuf.Timestamp
	2, // 4: odpf.assets.v1beta2.Experiment.Variant.attributes:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_odpf_assets_v1beta2_experiment_proto_init() }
func file_odpf_assets_v1beta2_experiment_proto_init() {
	if File_odpf_assets_v1beta2_experiment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_assets_v1beta2_experiment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experiment); i {
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
		file_odpf_assets_v1beta2_experiment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Experiment_Variant); i {
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
			RawDescriptor: file_odpf_assets_v1beta2_experiment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_assets_v1beta2_experiment_proto_goTypes,
		DependencyIndexes: file_odpf_assets_v1beta2_experiment_proto_depIdxs,
		MessageInfos:      file_odpf_assets_v1beta2_experiment_proto_msgTypes,
	}.Build()
	File_odpf_assets_v1beta2_experiment_proto = out.File
	file_odpf_assets_v1beta2_experiment_proto_rawDesc = nil
	file_odpf_assets_v1beta2_experiment_proto_goTypes = nil
	file_odpf_assets_v1beta2_experiment_proto_depIdxs = nil
}
