// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: types.proto

package proto

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

type Category int32

const (
	Category_CATEGORY_UNSPECIFIED              Category = 0 // zero value name should be suffixed with "_UNSPECIFIED".
	Category_CATEGORY_CLOTHING                 Category = 1 //should be prefixed with "CATEGORY_ // should be UPPER_SNAKE_CASE
	Category_CATEGORY_ELECTRONICS              Category = 2
	Category_CATEGORY_BOOKS                    Category = 3
	Category_CATEGORY_HOME_AND_KITCHEN         Category = 4
	Category_CATEGORY_SPORTS_AND_OUTDOORS      Category = 5
	Category_CATEGORY_BEAUTY_AND_PERSONAL_CARE Category = 6
	Category_CATEGORY_TOYS_AND_GAMES           Category = 7
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0: "CATEGORY_UNSPECIFIED",
		1: "CATEGORY_CLOTHING",
		2: "CATEGORY_ELECTRONICS",
		3: "CATEGORY_BOOKS",
		4: "CATEGORY_HOME_AND_KITCHEN",
		5: "CATEGORY_SPORTS_AND_OUTDOORS",
		6: "CATEGORY_BEAUTY_AND_PERSONAL_CARE",
		7: "CATEGORY_TOYS_AND_GAMES",
	}
	Category_value = map[string]int32{
		"CATEGORY_UNSPECIFIED":              0,
		"CATEGORY_CLOTHING":                 1,
		"CATEGORY_ELECTRONICS":              2,
		"CATEGORY_BOOKS":                    3,
		"CATEGORY_HOME_AND_KITCHEN":         4,
		"CATEGORY_SPORTS_AND_OUTDOORS":      5,
		"CATEGORY_BEAUTY_AND_PERSONAL_CARE": 6,
		"CATEGORY_TOYS_AND_GAMES":           7,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_types_proto_enumTypes[0].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_types_proto_enumTypes[0]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

// simple message
// PascalCase, such as "BlogRequest"
type BlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reserved 1; // uncomment this line to reserve the field numbers
	BlogId  int64  `protobuf:"varint,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"` //  lower_snake_case, such as "blog_id"
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BlogRequest) Reset() {
	*x = BlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogRequest) ProtoMessage() {}

func (x *BlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogRequest.ProtoReflect.Descriptor instead.
func (*BlogRequest) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *BlogRequest) GetBlogId() int64 {
	if x != nil {
		return x.BlogId
	}
	return 0
}

func (x *BlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string   `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Category  Category `protobuf:"varint,2,opt,name=category,proto3,enum=proto.Category" json:"category,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Product) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_CATEGORY_UNSPECIFIED
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2a, 0xee, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x4c,
	0x45, 0x43, 0x54, 0x52, 0x4f, 0x4e, 0x49, 0x43, 0x53, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x53, 0x10, 0x03, 0x12,
	0x1d, 0x0a, 0x19, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x48, 0x4f, 0x4d, 0x45,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4b, 0x49, 0x54, 0x43, 0x48, 0x45, 0x4e, 0x10, 0x04, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x53, 0x50, 0x4f, 0x52, 0x54,
	0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x44, 0x4f, 0x4f, 0x52, 0x53, 0x10, 0x05,
	0x12, 0x25, 0x0a, 0x21, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x42, 0x45, 0x41,
	0x55, 0x54, 0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c,
	0x5f, 0x43, 0x41, 0x52, 0x45, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x41, 0x54, 0x45, 0x47,
	0x4f, 0x52, 0x59, 0x5f, 0x54, 0x4f, 0x59, 0x53, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x47, 0x41, 0x4d,
	0x45, 0x53, 0x10, 0x07, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x62, 0x75,
	0x66, 0x2d, 0x62, 0x61, 0x73, 0x69, 0x63, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(Category)(0),       // 0: proto.Category
	(*BlogRequest)(nil), // 1: proto.BlogRequest
	(*Product)(nil),     // 2: proto.Product
}
var file_types_proto_depIdxs = []int32{
	0, // 0: proto.Product.category:type_name -> proto.Category
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogRequest); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		EnumInfos:         file_types_proto_enumTypes,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
