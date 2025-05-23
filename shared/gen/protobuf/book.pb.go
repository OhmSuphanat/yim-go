// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: shared/proto/book.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateBookRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	mi := &file_shared_proto_book_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PaginationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageId        int32                  `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_shared_proto_book_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{1}
}

func (x *PaginationRequest) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *PaginationRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetBookListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Pagination    *PaginationRequest     `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookListRequest) Reset() {
	*x = GetBookListRequest{}
	mi := &file_shared_proto_book_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListRequest) ProtoMessage() {}

func (x *GetBookListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListRequest.ProtoReflect.Descriptor instead.
func (*GetBookListRequest) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookListRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBookListRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_shared_proto_book_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{3}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Book) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Book) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PaginationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageId        int32                  `protobuf:"varint,1,opt,name=page_id,json=pageId,proto3" json:"page_id,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	LastPage      int32                  `protobuf:"varint,3,opt,name=last_page,json=lastPage,proto3" json:"last_page,omitempty"`
	TotalElements int32                  `protobuf:"varint,4,opt,name=total_elements,json=totalElements,proto3" json:"total_elements,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_shared_proto_book_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{4}
}

func (x *PaginationResponse) GetPageId() int32 {
	if x != nil {
		return x.PageId
	}
	return 0
}

func (x *PaginationResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationResponse) GetLastPage() int32 {
	if x != nil {
		return x.LastPage
	}
	return 0
}

func (x *PaginationResponse) GetTotalElements() int32 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

type GetBookListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Books         []*Book                `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	Pagination    *PaginationResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookListResponse) Reset() {
	*x = GetBookListResponse{}
	mi := &file_shared_proto_book_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookListResponse) ProtoMessage() {}

func (x *GetBookListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_book_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookListResponse.ProtoReflect.Descriptor instead.
func (*GetBookListResponse) Descriptor() ([]byte, []int) {
	return file_shared_proto_book_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookListResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *GetBookListResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_shared_proto_book_proto protoreflect.FileDescriptor

const file_shared_proto_book_proto_rawDesc = "" +
	"\n" +
	"\x17shared/proto/book.proto\x1a\x1bgoogle/protobuf/empty.proto\"c\n" +
	"\x11CreateBookRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x16\n" +
	"\x06author\x18\x02 \x01(\tR\x06author\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"I\n" +
	"\x11PaginationRequest\x12\x17\n" +
	"\apage_id\x18\x01 \x01(\x05R\x06pageId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\"X\n" +
	"\x12GetBookListRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x122\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x12.PaginationRequestR\n" +
	"pagination\"\xa4\x01\n" +
	"\x04Book\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x16\n" +
	"\x06author\x18\x03 \x01(\tR\x06author\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"created_at\x18\x05 \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\tR\tupdatedAt\"\x8e\x01\n" +
	"\x12PaginationResponse\x12\x17\n" +
	"\apage_id\x18\x01 \x01(\x05R\x06pageId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x1b\n" +
	"\tlast_page\x18\x03 \x01(\x05R\blastPage\x12%\n" +
	"\x0etotal_elements\x18\x04 \x01(\x05R\rtotalElements\"g\n" +
	"\x13GetBookListResponse\x12\x1b\n" +
	"\x05books\x18\x01 \x03(\v2\x05.BookR\x05books\x123\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2\x13.PaginationResponseR\n" +
	"pagination2\x81\x01\n" +
	"\vBookService\x128\n" +
	"\n" +
	"CreateBook\x12\x12.CreateBookRequest\x1a\x16.google.protobuf.Empty\x128\n" +
	"\vGetBookList\x12\x13.GetBookListRequest\x1a\x14.GetBookListResponseB\x15Z\x13shared/gen/protobufb\x06proto3"

var (
	file_shared_proto_book_proto_rawDescOnce sync.Once
	file_shared_proto_book_proto_rawDescData []byte
)

func file_shared_proto_book_proto_rawDescGZIP() []byte {
	file_shared_proto_book_proto_rawDescOnce.Do(func() {
		file_shared_proto_book_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shared_proto_book_proto_rawDesc), len(file_shared_proto_book_proto_rawDesc)))
	})
	return file_shared_proto_book_proto_rawDescData
}

var file_shared_proto_book_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shared_proto_book_proto_goTypes = []any{
	(*CreateBookRequest)(nil),   // 0: CreateBookRequest
	(*PaginationRequest)(nil),   // 1: PaginationRequest
	(*GetBookListRequest)(nil),  // 2: GetBookListRequest
	(*Book)(nil),                // 3: Book
	(*PaginationResponse)(nil),  // 4: PaginationResponse
	(*GetBookListResponse)(nil), // 5: GetBookListResponse
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_shared_proto_book_proto_depIdxs = []int32{
	1, // 0: GetBookListRequest.pagination:type_name -> PaginationRequest
	3, // 1: GetBookListResponse.books:type_name -> Book
	4, // 2: GetBookListResponse.pagination:type_name -> PaginationResponse
	0, // 3: BookService.CreateBook:input_type -> CreateBookRequest
	2, // 4: BookService.GetBookList:input_type -> GetBookListRequest
	6, // 5: BookService.CreateBook:output_type -> google.protobuf.Empty
	5, // 6: BookService.GetBookList:output_type -> GetBookListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_shared_proto_book_proto_init() }
func file_shared_proto_book_proto_init() {
	if File_shared_proto_book_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shared_proto_book_proto_rawDesc), len(file_shared_proto_book_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_proto_book_proto_goTypes,
		DependencyIndexes: file_shared_proto_book_proto_depIdxs,
		MessageInfos:      file_shared_proto_book_proto_msgTypes,
	}.Build()
	File_shared_proto_book_proto = out.File
	file_shared_proto_book_proto_goTypes = nil
	file_shared_proto_book_proto_depIdxs = nil
}
