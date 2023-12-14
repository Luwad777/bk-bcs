// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: content.proto

package pbcontent

import (
	base "github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/pkg/protocol/core/base"
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

// Content source resource reference: pkg/dal/table/content.go
type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec       *ContentSpec          `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Attachment *ContentAttachment    `protobuf:"bytes,3,opt,name=attachment,proto3" json:"attachment,omitempty"`
	Revision   *base.CreatedRevision `protobuf:"bytes,4,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetSpec() *ContentSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Content) GetAttachment() *ContentAttachment {
	if x != nil {
		return x.Attachment
	}
	return nil
}

func (x *Content) GetRevision() *base.CreatedRevision {
	if x != nil {
		return x.Revision
	}
	return nil
}

// ContentSpec source resource reference: pkg/dal/table/content.go
type ContentSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	ByteSize  uint64 `protobuf:"varint,2,opt,name=byte_size,json=byteSize,proto3" json:"byte_size,omitempty"`
}

func (x *ContentSpec) Reset() {
	*x = ContentSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentSpec) ProtoMessage() {}

func (x *ContentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentSpec.ProtoReflect.Descriptor instead.
func (*ContentSpec) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *ContentSpec) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ContentSpec) GetByteSize() uint64 {
	if x != nil {
		return x.ByteSize
	}
	return 0
}

// ReleasedContentSpec source resource reference: pkg/dal/table/content.go
type ReleasedContentSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature       string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	ByteSize        uint64 `protobuf:"varint,2,opt,name=byte_size,json=byteSize,proto3" json:"byte_size,omitempty"`
	OriginSignature string `protobuf:"bytes,3,opt,name=origin_signature,json=originSignature,proto3" json:"origin_signature,omitempty"`
	OriginByteSize  uint64 `protobuf:"varint,4,opt,name=origin_byte_size,json=originByteSize,proto3" json:"origin_byte_size,omitempty"`
}

func (x *ReleasedContentSpec) Reset() {
	*x = ReleasedContentSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReleasedContentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReleasedContentSpec) ProtoMessage() {}

func (x *ReleasedContentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReleasedContentSpec.ProtoReflect.Descriptor instead.
func (*ReleasedContentSpec) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

func (x *ReleasedContentSpec) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ReleasedContentSpec) GetByteSize() uint64 {
	if x != nil {
		return x.ByteSize
	}
	return 0
}

func (x *ReleasedContentSpec) GetOriginSignature() string {
	if x != nil {
		return x.OriginSignature
	}
	return ""
}

func (x *ReleasedContentSpec) GetOriginByteSize() uint64 {
	if x != nil {
		return x.OriginByteSize
	}
	return 0
}

// ContentAttachment source resource reference: pkg/dal/table/content.go
type ContentAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId        uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	AppId        uint32 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	ConfigItemId uint32 `protobuf:"varint,3,opt,name=config_item_id,json=configItemId,proto3" json:"config_item_id,omitempty"`
}

func (x *ContentAttachment) Reset() {
	*x = ContentAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentAttachment) ProtoMessage() {}

func (x *ContentAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentAttachment.ProtoReflect.Descriptor instead.
func (*ContentAttachment) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *ContentAttachment) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *ContentAttachment) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ContentAttachment) GetConfigItemId() uint32 {
	if x != nil {
		return x.ConfigItemId
	}
	return 0
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x79, 0x74,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x67, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x75, 0x65, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x62, 0x6b, 0x2d, 0x62, 0x63, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x63, 0x73, 0x2d, 0x62, 0x73, 0x63, 0x70, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3b, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_content_proto_goTypes = []interface{}{
	(*Content)(nil),              // 0: pbcontent.Content
	(*ContentSpec)(nil),          // 1: pbcontent.ContentSpec
	(*ReleasedContentSpec)(nil),  // 2: pbcontent.ReleasedContentSpec
	(*ContentAttachment)(nil),    // 3: pbcontent.ContentAttachment
	(*base.CreatedRevision)(nil), // 4: pbbase.CreatedRevision
}
var file_content_proto_depIdxs = []int32{
	1, // 0: pbcontent.Content.spec:type_name -> pbcontent.ContentSpec
	3, // 1: pbcontent.Content.attachment:type_name -> pbcontent.ContentAttachment
	4, // 2: pbcontent.Content.revision:type_name -> pbbase.CreatedRevision
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentSpec); i {
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
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReleasedContentSpec); i {
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
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentAttachment); i {
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
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
