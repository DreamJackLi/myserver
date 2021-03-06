// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: apiexport.proto

package apiexport

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	base "myserver/api/base"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExportReq struct {
	Head *base.Head `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	// 导出标签：企业 company；员工：employee
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	// 导出字段，以array包装后转为json字符串，格式实例:[{"field":"name","display":"企业名称"},{"field":"detailed_address","display":"详细地址"}]
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	// 公司归属省id
	ProvinceId int64 `protobuf:"varint,5,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	// 公司归属市id
	CityId int64 `protobuf:"varint,7,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	// 公司归属县/区id
	DistrictId int64 `protobuf:"varint,9,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	// 街道id
	StreetId int64 `protobuf:"varint,11,opt,name=street_id,json=streetId,proto3" json:"street_id,omitempty"`
	// 社区id
	CommunityId int64 `protobuf:"varint,13,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
}

func (m *ExportReq) Reset()         { *m = ExportReq{} }
func (m *ExportReq) String() string { return proto.CompactTextString(m) }
func (*ExportReq) ProtoMessage()    {}
func (*ExportReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4335b8daeb8c6233, []int{0}
}
func (m *ExportReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportReq.Merge(m, src)
}
func (m *ExportReq) XXX_Size() int {
	return m.Size()
}
func (m *ExportReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportReq.DiscardUnknown(m)
}

var xxx_messageInfo_ExportReq proto.InternalMessageInfo

func (m *ExportReq) GetHead() *base.Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *ExportReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ExportReq) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ExportReq) GetProvinceId() int64 {
	if m != nil {
		return m.ProvinceId
	}
	return 0
}

func (m *ExportReq) GetCityId() int64 {
	if m != nil {
		return m.CityId
	}
	return 0
}

func (m *ExportReq) GetDistrictId() int64 {
	if m != nil {
		return m.DistrictId
	}
	return 0
}

func (m *ExportReq) GetStreetId() int64 {
	if m != nil {
		return m.StreetId
	}
	return 0
}

func (m *ExportReq) GetCommunityId() int64 {
	if m != nil {
		return m.CommunityId
	}
	return 0
}

type ExportRes struct {
	Base *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// 文件包地址
	FileAddress string `protobuf:"bytes,2,opt,name=file_address,json=fileAddress,proto3" json:"file_address,omitempty"`
}

func (m *ExportRes) Reset()         { *m = ExportRes{} }
func (m *ExportRes) String() string { return proto.CompactTextString(m) }
func (*ExportRes) ProtoMessage()    {}
func (*ExportRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4335b8daeb8c6233, []int{1}
}
func (m *ExportRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportRes.Merge(m, src)
}
func (m *ExportRes) XXX_Size() int {
	return m.Size()
}
func (m *ExportRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportRes.DiscardUnknown(m)
}

var xxx_messageInfo_ExportRes proto.InternalMessageInfo

func (m *ExportRes) GetBase() *base.BaseRes {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *ExportRes) GetFileAddress() string {
	if m != nil {
		return m.FileAddress
	}
	return ""
}

type ExportRecord struct {
	// 主键
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 导出标签
	Tag string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	// 统计sql
	CountSql string `protobuf:"bytes,5,opt,name=count_sql,json=countSql,proto3" json:"count_sql,omitempty"`
	// 查询sql
	QuerySql string `protobuf:"bytes,6,opt,name=query_sql,json=querySql,proto3" json:"query_sql,omitempty"`
	// 文件包地址
	FilePath string `protobuf:"bytes,7,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// 状态 1.创建任务；2.数据写入文件；3文件上传；4.完成
	Status int32 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *ExportRecord) Reset()         { *m = ExportRecord{} }
func (m *ExportRecord) String() string { return proto.CompactTextString(m) }
func (*ExportRecord) ProtoMessage()    {}
func (*ExportRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_4335b8daeb8c6233, []int{2}
}
func (m *ExportRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExportRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExportRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExportRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportRecord.Merge(m, src)
}
func (m *ExportRecord) XXX_Size() int {
	return m.Size()
}
func (m *ExportRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ExportRecord proto.InternalMessageInfo

func (m *ExportRecord) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExportRecord) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ExportRecord) GetCountSql() string {
	if m != nil {
		return m.CountSql
	}
	return ""
}

func (m *ExportRecord) GetQuerySql() string {
	if m != nil {
		return m.QuerySql
	}
	return ""
}

func (m *ExportRecord) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *ExportRecord) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Field struct {
	// 字段名
	Filed string `protobuf:"bytes,1,opt,name=filed,proto3" json:"filed,omitempty"`
	// 页面展示名称
	Display string `protobuf:"bytes,2,opt,name=display,proto3" json:"display,omitempty"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_4335b8daeb8c6233, []int{3}
}
func (m *Field) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Field.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return m.Size()
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetFiled() string {
	if m != nil {
		return m.Filed
	}
	return ""
}

func (m *Field) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func init() {
	proto.RegisterType((*ExportReq)(nil), "apiexport.ExportReq")
	proto.RegisterType((*ExportRes)(nil), "apiexport.ExportRes")
	proto.RegisterType((*ExportRecord)(nil), "apiexport.ExportRecord")
	proto.RegisterType((*Field)(nil), "apiexport.Field")
}

func init() { proto.RegisterFile("apiexport.proto", fileDescriptor_4335b8daeb8c6233) }

var fileDescriptor_4335b8daeb8c6233 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xe3, 0xa4, 0xd9, 0xc4, 0xb3, 0x2d, 0x20, 0xab, 0x82, 0x55, 0x8b, 0x96, 0x34, 0xa7,
	0x9c, 0x52, 0xa9, 0x20, 0x71, 0x43, 0xa2, 0x12, 0x88, 0xbd, 0x81, 0x79, 0x80, 0xc8, 0x5d, 0x4f,
	0x89, 0xa5, 0x4d, 0x76, 0x63, 0x3b, 0x15, 0x79, 0x0b, 0xde, 0x80, 0xd7, 0xe1, 0xd8, 0x23, 0x47,
	0x94, 0xdc, 0x79, 0x86, 0xca, 0xe3, 0x4d, 0x7a, 0xe9, 0x65, 0xb5, 0xf3, 0xfd, 0xe3, 0xd1, 0xfc,
	0x33, 0x03, 0xcf, 0x55, 0x63, 0xf0, 0x67, 0x53, 0x5b, 0x3f, 0x6d, 0x6c, 0xed, 0x6b, 0xc1, 0x0f,
	0xe0, 0xec, 0xf5, 0x62, 0xe3, 0xd0, 0xde, 0xa1, 0xbd, 0x54, 0x8d, 0xb9, 0xbc, 0x51, 0x0e, 0xe9,
	0x13, 0x13, 0xc7, 0xff, 0x19, 0xf0, 0x4f, 0x94, 0x28, 0x71, 0x25, 0x72, 0x38, 0x9a, 0xa3, 0xd2,
	0x19, 0x1b, 0xb1, 0x49, 0x7a, 0x05, 0x53, 0x4a, 0xfc, 0x82, 0x4a, 0x4b, 0xe2, 0xe2, 0x05, 0xf4,
	0xbc, 0xfa, 0x91, 0x75, 0x47, 0x6c, 0xc2, 0x65, 0xf8, 0x15, 0xa7, 0xd0, 0xbf, 0x35, 0x58, 0xe9,
	0xac, 0x47, 0x2c, 0x06, 0xe2, 0x0d, 0xa4, 0x8d, 0xad, 0xef, 0xcc, 0xb2, 0xc4, 0x99, 0xd1, 0x59,
	0x7f, 0xc4, 0x26, 0x3d, 0x09, 0x7b, 0x54, 0x68, 0xf1, 0x0a, 0x06, 0xa5, 0xf1, 0x9b, 0x20, 0x0e,
	0x48, 0x4c, 0x42, 0x58, 0xd0, 0x4b, 0x6d, 0x9c, 0xb7, 0xa6, 0xf4, 0x41, 0xe4, 0xf1, 0xe5, 0x1e,
	0x15, 0x5a, 0x9c, 0x03, 0x77, 0xde, 0x22, 0x92, 0x9c, 0x92, 0x3c, 0x8c, 0xa0, 0xd0, 0xe2, 0x02,
	0x8e, 0xcb, 0x7a, 0xb1, 0x58, 0x2f, 0xdb, 0xda, 0x27, 0xa4, 0xa7, 0x07, 0x56, 0xe8, 0xf1, 0xb7,
	0x47, 0xbf, 0x4e, 0x5c, 0xc0, 0x51, 0xb0, 0xd8, 0xfa, 0x3d, 0x89, 0x7e, 0xaf, 0x95, 0x43, 0x89,
	0x4e, 0x92, 0x14, 0x4a, 0xde, 0x9a, 0x0a, 0x67, 0x4a, 0x6b, 0x8b, 0xce, 0xb5, 0xde, 0xd3, 0xc0,
	0x3e, 0x46, 0x34, 0xfe, 0xcd, 0xe0, 0x78, 0x5f, 0xb3, 0xac, 0xad, 0x16, 0xcf, 0xa0, 0x6b, 0xe2,
	0x10, 0x7b, 0xb2, 0x6b, 0x9e, 0x1a, 0xdb, 0x39, 0xf0, 0xb2, 0x5e, 0x2f, 0xfd, 0xcc, 0xad, 0x2a,
	0x1a, 0x0f, 0x97, 0x43, 0x02, 0xdf, 0x57, 0x55, 0x10, 0x57, 0x6b, 0xb4, 0x1b, 0x12, 0x93, 0x28,
	0x12, 0x68, 0x45, 0xea, 0xa7, 0x51, 0x7e, 0x4e, 0xb3, 0xe3, 0x72, 0x18, 0xc0, 0x57, 0xe5, 0xe7,
	0xe2, 0x25, 0x24, 0xce, 0x2b, 0xbf, 0x76, 0xd9, 0x70, 0xc4, 0x26, 0x7d, 0xd9, 0x46, 0xe3, 0xf7,
	0xd0, 0xff, 0x4c, 0x8b, 0xa1, 0x75, 0x55, 0x18, 0x9b, 0xa3, 0x75, 0x55, 0xa8, 0x45, 0x06, 0x03,
	0x6d, 0x5c, 0x53, 0xa9, 0x4d, 0xdb, 0xe3, 0x3e, 0xbc, 0xfa, 0x00, 0x49, 0x74, 0x26, 0xde, 0x1d,
	0xfe, 0x4e, 0xa7, 0x8f, 0xd7, 0x76, 0x38, 0x9d, 0xb3, 0xa7, 0xa8, 0xbb, 0xce, 0xfe, 0x6c, 0x73,
	0x76, 0xbf, 0xcd, 0xd9, 0xbf, 0x6d, 0xce, 0x7e, 0xed, 0xf2, 0xce, 0xfd, 0x2e, 0xef, 0xfc, 0xdd,
	0xe5, 0x9d, 0x9b, 0x84, 0xee, 0xef, 0xed, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xed, 0x12,
	0x43, 0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExportClient is the client API for Export service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExportClient interface {
	// 导出通用方法
	Export(ctx context.Context, in *ExportReq, opts ...grpc.CallOption) (*ExportRes, error)
}

type exportClient struct {
	cc *grpc.ClientConn
}

func NewExportClient(cc *grpc.ClientConn) ExportClient {
	return &exportClient{cc}
}

func (c *exportClient) Export(ctx context.Context, in *ExportReq, opts ...grpc.CallOption) (*ExportRes, error) {
	out := new(ExportRes)
	err := c.cc.Invoke(ctx, "/apiexport.Export/Export", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExportServer is the server API for Export service.
type ExportServer interface {
	// 导出通用方法
	Export(context.Context, *ExportReq) (*ExportRes, error)
}

// UnimplementedExportServer can be embedded to have forward compatible implementations.
type UnimplementedExportServer struct {
}

func (*UnimplementedExportServer) Export(ctx context.Context, req *ExportReq) (*ExportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Export not implemented")
}

func RegisterExportServer(s *grpc.Server, srv ExportServer) {
	s.RegisterService(&_Export_serviceDesc, srv)
}

func _Export_Export_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).Export(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiexport.Export/Export",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).Export(ctx, req.(*ExportReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Export_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apiexport.Export",
	HandlerType: (*ExportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Export",
			Handler:    _Export_Export_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apiexport.proto",
}

func (m *ExportReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CommunityId != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.CommunityId))
		i--
		dAtA[i] = 0x68
	}
	if m.StreetId != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.StreetId))
		i--
		dAtA[i] = 0x58
	}
	if m.DistrictId != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.DistrictId))
		i--
		dAtA[i] = 0x48
	}
	if m.CityId != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.CityId))
		i--
		dAtA[i] = 0x38
	}
	if m.ProvinceId != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.ProvinceId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Field) > 0 {
		i -= len(m.Field)
		copy(dAtA[i:], m.Field)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.Field)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x12
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiexport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExportRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FileAddress) > 0 {
		i -= len(m.FileAddress)
		copy(dAtA[i:], m.FileAddress)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.FileAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Base != nil {
		{
			size, err := m.Base.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApiexport(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExportRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExportRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExportRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.FilePath) > 0 {
		i -= len(m.FilePath)
		copy(dAtA[i:], m.FilePath)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.FilePath)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.QuerySql) > 0 {
		i -= len(m.QuerySql)
		copy(dAtA[i:], m.QuerySql)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.QuerySql)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CountSql) > 0 {
		i -= len(m.CountSql)
		copy(dAtA[i:], m.CountSql)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.CountSql)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApiexport(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Field) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Field) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Field) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Display) > 0 {
		i -= len(m.Display)
		copy(dAtA[i:], m.Display)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.Display)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Filed) > 0 {
		i -= len(m.Filed)
		copy(dAtA[i:], m.Filed)
		i = encodeVarintApiexport(dAtA, i, uint64(len(m.Filed)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApiexport(dAtA []byte, offset int, v uint64) int {
	offset -= sovApiexport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExportReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.Field)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	if m.ProvinceId != 0 {
		n += 1 + sovApiexport(uint64(m.ProvinceId))
	}
	if m.CityId != 0 {
		n += 1 + sovApiexport(uint64(m.CityId))
	}
	if m.DistrictId != 0 {
		n += 1 + sovApiexport(uint64(m.DistrictId))
	}
	if m.StreetId != 0 {
		n += 1 + sovApiexport(uint64(m.StreetId))
	}
	if m.CommunityId != 0 {
		n += 1 + sovApiexport(uint64(m.CommunityId))
	}
	return n
}

func (m *ExportRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Base != nil {
		l = m.Base.Size()
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.FileAddress)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	return n
}

func (m *ExportRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApiexport(uint64(m.Id))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.CountSql)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.QuerySql)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.FilePath)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovApiexport(uint64(m.Status))
	}
	return n
}

func (m *Field) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filed)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	l = len(m.Display)
	if l > 0 {
		n += 1 + l + sovApiexport(uint64(l))
	}
	return n
}

func sovApiexport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApiexport(x uint64) (n int) {
	return sovApiexport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExportReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiexport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &base.Head{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Field", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Field = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProvinceId", wireType)
			}
			m.ProvinceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProvinceId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CityId", wireType)
			}
			m.CityId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CityId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrictId", wireType)
			}
			m.DistrictId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DistrictId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StreetId", wireType)
			}
			m.StreetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StreetId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityId", wireType)
			}
			m.CommunityId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommunityId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApiexport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiexport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Base == nil {
				m.Base = &base.BaseRes{}
			}
			if err := m.Base.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiexport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExportRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiexport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExportRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExportRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountSql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CountSql = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuerySql", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuerySql = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilePath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FilePath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApiexport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Field) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApiexport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Field: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Field: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Display", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApiexport
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApiexport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Display = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApiexport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApiexport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApiexport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApiexport
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApiexport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApiexport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApiexport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApiexport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApiexport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApiexport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApiexport = fmt.Errorf("proto: unexpected end of group")
)
