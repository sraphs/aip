// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: freight/v1/site.proto

package freight

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A site is a node in a [shipper][freight.v1.Shipper]'s
// transport network.
type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the site.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The creation timestamp of the site.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update timestamp of the site.
	//
	// Updated when create/update/delete operation is performed.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The deletion timestamp of the site.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// The display name of the site.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The geographic location of the site.
	LatLng *latlng.LatLng `protobuf:"bytes,6,opt,name=lat_lng,json=latLng,proto3" json:"lat_lng,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_freight_v1_site_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_freight_v1_site_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_freight_v1_site_proto_rawDescGZIP(), []int{0}
}

func (x *Site) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Site) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Site) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Site) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *Site) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Site) GetLatLng() *latlng.LatLng {
	if x != nil {
		return x.LatLng
	}
	return nil
}

var File_freight_v1_site_proto protoreflect.FileDescriptor

var file_freight_v1_site_proto_rawDesc = []byte{
	0x0a, 0x15, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61,
	0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x03, 0x0a, 0x04, 0x53,
	0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x5f, 0x6c, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74,
	0x4c, 0x6e, 0x67, 0x3a, 0x49, 0xea, 0x41, 0x46, 0x0a, 0x16, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x2e, 0x73, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x69, 0x74, 0x65,
	0x12, 0x1f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x65, 0x72, 0x7d, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x69, 0x74, 0x65,
	0x7d, 0x2a, 0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x32, 0x04, 0x73, 0x69, 0x74, 0x65, 0x42, 0x3c,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x61,
	0x70, 0x68, 0x73, 0x2f, 0x61, 0x69, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freight_v1_site_proto_rawDescOnce sync.Once
	file_freight_v1_site_proto_rawDescData = file_freight_v1_site_proto_rawDesc
)

func file_freight_v1_site_proto_rawDescGZIP() []byte {
	file_freight_v1_site_proto_rawDescOnce.Do(func() {
		file_freight_v1_site_proto_rawDescData = protoimpl.X.CompressGZIP(file_freight_v1_site_proto_rawDescData)
	})
	return file_freight_v1_site_proto_rawDescData
}

var file_freight_v1_site_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_freight_v1_site_proto_goTypes = []interface{}{
	(*Site)(nil),                  // 0: freight.v1.Site
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*latlng.LatLng)(nil),         // 2: google.type.LatLng
}
var file_freight_v1_site_proto_depIdxs = []int32{
	1, // 0: freight.v1.Site.create_time:type_name -> google.protobuf.Timestamp
	1, // 1: freight.v1.Site.update_time:type_name -> google.protobuf.Timestamp
	1, // 2: freight.v1.Site.delete_time:type_name -> google.protobuf.Timestamp
	2, // 3: freight.v1.Site.lat_lng:type_name -> google.type.LatLng
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_freight_v1_site_proto_init() }
func file_freight_v1_site_proto_init() {
	if File_freight_v1_site_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_freight_v1_site_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
			RawDescriptor: file_freight_v1_site_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freight_v1_site_proto_goTypes,
		DependencyIndexes: file_freight_v1_site_proto_depIdxs,
		MessageInfos:      file_freight_v1_site_proto_msgTypes,
	}.Build()
	File_freight_v1_site_proto = out.File
	file_freight_v1_site_proto_rawDesc = nil
	file_freight_v1_site_proto_goTypes = nil
	file_freight_v1_site_proto_depIdxs = nil
}
