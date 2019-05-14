// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/todos/todos.proto

package go_micro_api_todos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Todo struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{1}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type AddReq struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddReq) Reset()         { *m = AddReq{} }
func (m *AddReq) String() string { return proto.CompactTextString(m) }
func (*AddReq) ProtoMessage()    {}
func (*AddReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{3}
}

func (m *AddReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddReq.Unmarshal(m, b)
}
func (m *AddReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddReq.Marshal(b, m, deterministic)
}
func (m *AddReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddReq.Merge(m, src)
}
func (m *AddReq) XXX_Size() int {
	return xxx_messageInfo_AddReq.Size(m)
}
func (m *AddReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddReq proto.InternalMessageInfo

func (m *AddReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type AddResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResp) Reset()         { *m = AddResp{} }
func (m *AddResp) String() string { return proto.CompactTextString(m) }
func (*AddResp) ProtoMessage()    {}
func (*AddResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{4}
}

func (m *AddResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResp.Unmarshal(m, b)
}
func (m *AddResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResp.Marshal(b, m, deterministic)
}
func (m *AddResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResp.Merge(m, src)
}
func (m *AddResp) XXX_Size() int {
	return xxx_messageInfo_AddResp.Size(m)
}
func (m *AddResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddResp proto.InternalMessageInfo

func (m *AddResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AddResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type DelReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelReq) Reset()         { *m = DelReq{} }
func (m *DelReq) String() string { return proto.CompactTextString(m) }
func (*DelReq) ProtoMessage()    {}
func (*DelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{5}
}

func (m *DelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelReq.Unmarshal(m, b)
}
func (m *DelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelReq.Marshal(b, m, deterministic)
}
func (m *DelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelReq.Merge(m, src)
}
func (m *DelReq) XXX_Size() int {
	return xxx_messageInfo_DelReq.Size(m)
}
func (m *DelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelReq proto.InternalMessageInfo

func (m *DelReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DelResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelResp) Reset()         { *m = DelResp{} }
func (m *DelResp) String() string { return proto.CompactTextString(m) }
func (*DelResp) ProtoMessage()    {}
func (*DelResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{6}
}

func (m *DelResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelResp.Unmarshal(m, b)
}
func (m *DelResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelResp.Marshal(b, m, deterministic)
}
func (m *DelResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelResp.Merge(m, src)
}
func (m *DelResp) XXX_Size() int {
	return xxx_messageInfo_DelResp.Size(m)
}
func (m *DelResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelResp proto.InternalMessageInfo

func (m *DelResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DelResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type EditReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditReq) Reset()         { *m = EditReq{} }
func (m *EditReq) String() string { return proto.CompactTextString(m) }
func (*EditReq) ProtoMessage()    {}
func (*EditReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{7}
}

func (m *EditReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditReq.Unmarshal(m, b)
}
func (m *EditReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditReq.Marshal(b, m, deterministic)
}
func (m *EditReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditReq.Merge(m, src)
}
func (m *EditReq) XXX_Size() int {
	return xxx_messageInfo_EditReq.Size(m)
}
func (m *EditReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditReq proto.InternalMessageInfo

func (m *EditReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EditReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EditReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type EditResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditResp) Reset()         { *m = EditResp{} }
func (m *EditResp) String() string { return proto.CompactTextString(m) }
func (*EditResp) ProtoMessage()    {}
func (*EditResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{8}
}

func (m *EditResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditResp.Unmarshal(m, b)
}
func (m *EditResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditResp.Marshal(b, m, deterministic)
}
func (m *EditResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditResp.Merge(m, src)
}
func (m *EditResp) XXX_Size() int {
	return xxx_messageInfo_EditResp.Size(m)
}
func (m *EditResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EditResp.DiscardUnknown(m)
}

var xxx_messageInfo_EditResp proto.InternalMessageInfo

func (m *EditResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EditResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ListReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListReq) Reset()         { *m = ListReq{} }
func (m *ListReq) String() string { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()    {}
func (*ListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{9}
}

func (m *ListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReq.Unmarshal(m, b)
}
func (m *ListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReq.Marshal(b, m, deterministic)
}
func (m *ListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReq.Merge(m, src)
}
func (m *ListReq) XXX_Size() int {
	return xxx_messageInfo_ListReq.Size(m)
}
func (m *ListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListReq proto.InternalMessageInfo

type ListResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Todos                []*Todo  `protobuf:"bytes,3,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResp) Reset()         { *m = ListResp{} }
func (m *ListResp) String() string { return proto.CompactTextString(m) }
func (*ListResp) ProtoMessage()    {}
func (*ListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{10}
}

func (m *ListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResp.Unmarshal(m, b)
}
func (m *ListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResp.Marshal(b, m, deterministic)
}
func (m *ListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResp.Merge(m, src)
}
func (m *ListResp) XXX_Size() int {
	return xxx_messageInfo_ListResp.Size(m)
}
func (m *ListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListResp proto.InternalMessageInfo

func (m *ListResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ListResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ListResp) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type DetailReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailReq) Reset()         { *m = DetailReq{} }
func (m *DetailReq) String() string { return proto.CompactTextString(m) }
func (*DetailReq) ProtoMessage()    {}
func (*DetailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{11}
}

func (m *DetailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailReq.Unmarshal(m, b)
}
func (m *DetailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailReq.Marshal(b, m, deterministic)
}
func (m *DetailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailReq.Merge(m, src)
}
func (m *DetailReq) XXX_Size() int {
	return xxx_messageInfo_DetailReq.Size(m)
}
func (m *DetailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailReq.DiscardUnknown(m)
}

var xxx_messageInfo_DetailReq proto.InternalMessageInfo

func (m *DetailReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DetailResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Todo                 *Todo    `protobuf:"bytes,3,opt,name=todo,proto3" json:"todo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetailResp) Reset()         { *m = DetailResp{} }
func (m *DetailResp) String() string { return proto.CompactTextString(m) }
func (*DetailResp) ProtoMessage()    {}
func (*DetailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1dfee0949634faa, []int{12}
}

func (m *DetailResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailResp.Unmarshal(m, b)
}
func (m *DetailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailResp.Marshal(b, m, deterministic)
}
func (m *DetailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailResp.Merge(m, src)
}
func (m *DetailResp) XXX_Size() int {
	return xxx_messageInfo_DetailResp.Size(m)
}
func (m *DetailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailResp.DiscardUnknown(m)
}

var xxx_messageInfo_DetailResp proto.InternalMessageInfo

func (m *DetailResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DetailResp) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *DetailResp) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.api.todos.Message")
	proto.RegisterType((*Todo)(nil), "go.micro.api.todos.Todo")
	proto.RegisterType((*Error)(nil), "go.micro.api.todos.Error")
	proto.RegisterType((*AddReq)(nil), "go.micro.api.todos.AddReq")
	proto.RegisterType((*AddResp)(nil), "go.micro.api.todos.AddResp")
	proto.RegisterType((*DelReq)(nil), "go.micro.api.todos.DelReq")
	proto.RegisterType((*DelResp)(nil), "go.micro.api.todos.DelResp")
	proto.RegisterType((*EditReq)(nil), "go.micro.api.todos.EditReq")
	proto.RegisterType((*EditResp)(nil), "go.micro.api.todos.EditResp")
	proto.RegisterType((*ListReq)(nil), "go.micro.api.todos.ListReq")
	proto.RegisterType((*ListResp)(nil), "go.micro.api.todos.ListResp")
	proto.RegisterType((*DetailReq)(nil), "go.micro.api.todos.DetailReq")
	proto.RegisterType((*DetailResp)(nil), "go.micro.api.todos.DetailResp")
}

func init() { proto.RegisterFile("proto/todos/todos.proto", fileDescriptor_c1dfee0949634faa) }

var fileDescriptor_c1dfee0949634faa = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0x6e, 0xfe, 0xdb, 0x53, 0x10, 0x39, 0x88, 0xc6, 0x54, 0xa5, 0xcc, 0xaa, 0x0b, 0x49, 0xa1,
	0x05, 0x71, 0x69, 0xb1, 0x5d, 0x14, 0x74, 0x13, 0xea, 0x03, 0xc4, 0xcc, 0x50, 0x06, 0x62, 0x27,
	0x66, 0xc6, 0x85, 0x0f, 0xd0, 0xf7, 0xf4, 0x51, 0x64, 0x7e, 0xc2, 0xe5, 0x72, 0xa7, 0x85, 0x0b,
	0xd9, 0x94, 0x39, 0x3f, 0xdf, 0x77, 0x3e, 0xce, 0xf9, 0x1a, 0x78, 0xd3, 0xf5, 0x42, 0x89, 0xb5,
	0x12, 0x54, 0x48, 0xfb, 0x5b, 0x9a, 0x0c, 0xe2, 0x59, 0x94, 0xbf, 0x78, 0xd3, 0x8b, 0xb2, 0xee,
	0x78, 0x69, 0x2a, 0x64, 0x01, 0xd9, 0x77, 0x26, 0x65, 0x7d, 0x66, 0xf8, 0x12, 0x22, 0x59, 0xff,
	0xcd, 0x83, 0x65, 0xb0, 0x9a, 0x55, 0xfa, 0x49, 0x3e, 0x41, 0x7c, 0x12, 0x54, 0xe0, 0x2b, 0x48,
	0x14, 0x57, 0x2d, 0x73, 0x35, 0x1b, 0x60, 0x0e, 0x59, 0x23, 0x2e, 0x8a, 0x5d, 0x54, 0x1e, 0x9a,
	0xfc, 0x10, 0x92, 0x2d, 0x24, 0x87, 0xbe, 0x17, 0x3d, 0x22, 0xc4, 0x8d, 0xa0, 0x16, 0x97, 0x54,
	0xe6, 0x8d, 0xaf, 0x21, 0xa5, 0x4c, 0xd5, 0xbc, 0x75, 0x28, 0x17, 0x91, 0xcf, 0x90, 0xee, 0x28,
	0xad, 0xd8, 0xef, 0x67, 0x8f, 0x3b, 0x41, 0x66, 0x90, 0xb2, 0xd3, 0x4d, 0xf2, 0x4f, 0xd3, 0x30,
	0x29, 0x0d, 0x78, 0x5a, 0x0d, 0x21, 0xae, 0x21, 0x61, 0x5a, 0x93, 0x01, 0xcf, 0x37, 0x6f, 0xcb,
	0xa7, 0xcb, 0x28, 0x8d, 0xe8, 0xca, 0xf6, 0x91, 0x1c, 0xd2, 0x3d, 0x6b, 0xb5, 0x9e, 0x17, 0x10,
	0x72, 0xea, 0xc4, 0x84, 0x9c, 0xea, 0x79, 0xa6, 0x32, 0xee, 0xbc, 0x23, 0x64, 0x07, 0xca, 0x95,
	0x67, 0xe0, 0xc3, 0x42, 0xc2, 0x1b, 0x0b, 0x89, 0x1e, 0x2f, 0xe4, 0x07, 0x4c, 0x2d, 0xd5, 0xb8,
	0x0a, 0x67, 0x90, 0x7d, 0xe3, 0x52, 0x2b, 0x24, 0xd7, 0x00, 0xa6, 0xf6, 0x3d, 0xea, 0x08, 0x2c,
	0x21, 0x31, 0xd9, 0x3c, 0x5a, 0x46, 0xab, 0xf9, 0x26, 0xf7, 0x01, 0xb4, 0x25, 0xab, 0x64, 0xb0,
	0xef, 0x6c, 0x6f, 0xec, 0xe3, 0xbb, 0xd3, 0x35, 0x00, 0x18, 0xaa, 0xe3, 0xca, 0xfc, 0x08, 0xb1,
	0xce, 0x9a, 0xbd, 0xdf, 0x53, 0x69, 0xba, 0x36, 0xff, 0x42, 0x48, 0x74, 0x28, 0xf1, 0x0b, 0x44,
	0x3b, 0x4a, 0xb1, 0xf0, 0x01, 0xac, 0xf9, 0x8b, 0xc5, 0xcd, 0x9a, 0xec, 0xc8, 0x44, 0x33, 0xec,
	0x59, 0xeb, 0x67, 0xb0, 0x76, 0xf5, 0x33, 0x38, 0xc3, 0x92, 0x09, 0x7e, 0x85, 0x58, 0x9b, 0x03,
	0xbd, 0x6d, 0xce, 0x81, 0xc5, 0xbb, 0xdb, 0xc5, 0x81, 0x44, 0x9f, 0xdf, 0x4f, 0xe2, 0x4c, 0xe2,
	0x27, 0x19, 0x5c, 0x43, 0x26, 0x78, 0xd4, 0xff, 0x30, 0x7d, 0x1e, 0x7c, 0xef, 0x97, 0xec, 0x0e,
	0x5b, 0x7c, 0xb8, 0x57, 0xd6, 0x54, 0x3f, 0x53, 0xf3, 0x85, 0xdb, 0xfe, 0x0f, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x0f, 0x92, 0x04, 0xfc, 0x04, 0x00, 0x00,
}
