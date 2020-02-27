// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

package base

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

type Api int32

const (
	Api_Empty                                  Api = 0
	Api_BuildConn                              Api = 1
	Api_WebLogin                               Api = 2
	Api_AddEmployeeInfo                        Api = 3
	Api_UpdateEmployeeInfo                     Api = 4
	Api_GetEmployeeInfo                        Api = 5
	Api_AddEmployeeHealthInfo                  Api = 6
	Api_UpdateEmployeeHealthInfo               Api = 7
	Api_GetEmployeeHealthInfo                  Api = 8
	Api_AddEmployeeHealthRecord                Api = 9
	Api_AddCompanyInfo                         Api = 10
	Api_UpdateCompanyInfo                      Api = 11
	Api_GetCompanyInfo                         Api = 12
	Api_AddPreventionRecord                    Api = 13
	Api_UpdatePreventionRecord                 Api = 14
	Api_GetPreventionRecord                    Api = 15
	Api_GetQiniuToken                          Api = 16
	Api_GetProvinces                           Api = 17
	Api_GetCity                                Api = 18
	Api_GetDistrict                            Api = 19
	Api_GetStreet                              Api = 20
	Api_GetCommunity                           Api = 21
	Api_QueryEmployInfoByUser                  Api = 22
	Api_GetCompanyInfoList                     Api = 23
	Api_DeleteCompanyInfo                      Api = 24
	Api_GetPreventionRecordList                Api = 25
	Api_GetEmployeeHealthRecordList            Api = 26
	Api_DeleteEmployInfo                       Api = 27
	Api_GetEmployeeHealthRecordListByCompanyId Api = 28
)

var Api_name = map[int32]string{
	0:  "Empty",
	1:  "BuildConn",
	2:  "WebLogin",
	3:  "AddEmployeeInfo",
	4:  "UpdateEmployeeInfo",
	5:  "GetEmployeeInfo",
	6:  "AddEmployeeHealthInfo",
	7:  "UpdateEmployeeHealthInfo",
	8:  "GetEmployeeHealthInfo",
	9:  "AddEmployeeHealthRecord",
	10: "AddCompanyInfo",
	11: "UpdateCompanyInfo",
	12: "GetCompanyInfo",
	13: "AddPreventionRecord",
	14: "UpdatePreventionRecord",
	15: "GetPreventionRecord",
	16: "GetQiniuToken",
	17: "GetProvinces",
	18: "GetCity",
	19: "GetDistrict",
	20: "GetStreet",
	21: "GetCommunity",
	22: "QueryEmployInfoByUser",
	23: "GetCompanyInfoList",
	24: "DeleteCompanyInfo",
	25: "GetPreventionRecordList",
	26: "GetEmployeeHealthRecordList",
	27: "DeleteEmployInfo",
	28: "GetEmployeeHealthRecordListByCompanyId",
}

var Api_value = map[string]int32{
	"Empty":                                  0,
	"BuildConn":                              1,
	"WebLogin":                               2,
	"AddEmployeeInfo":                        3,
	"UpdateEmployeeInfo":                     4,
	"GetEmployeeInfo":                        5,
	"AddEmployeeHealthInfo":                  6,
	"UpdateEmployeeHealthInfo":               7,
	"GetEmployeeHealthInfo":                  8,
	"AddEmployeeHealthRecord":                9,
	"AddCompanyInfo":                         10,
	"UpdateCompanyInfo":                      11,
	"GetCompanyInfo":                         12,
	"AddPreventionRecord":                    13,
	"UpdatePreventionRecord":                 14,
	"GetPreventionRecord":                    15,
	"GetQiniuToken":                          16,
	"GetProvinces":                           17,
	"GetCity":                                18,
	"GetDistrict":                            19,
	"GetStreet":                              20,
	"GetCommunity":                           21,
	"QueryEmployInfoByUser":                  22,
	"GetCompanyInfoList":                     23,
	"DeleteCompanyInfo":                      24,
	"GetPreventionRecordList":                25,
	"GetEmployeeHealthRecordList":            26,
	"DeleteEmployInfo":                       27,
	"GetEmployeeHealthRecordListByCompanyId": 28,
}

func (x Api) String() string {
	return proto.EnumName(Api_name, int32(x))
}

func (Api) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

type Head struct {
	// int32 ApiID = 1;
	ConnID   string `protobuf:"bytes,1,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
	GroupID  string `protobuf:"bytes,2,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	AppKey   string `protobuf:"bytes,3,opt,name=AppKey,proto3" json:"AppKey,omitempty"`
	ServerID string `protobuf:"bytes,4,opt,name=ServerID,proto3" json:"ServerID,omitempty"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}
func (m *Head) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Head.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head.Merge(m, src)
}
func (m *Head) XXX_Size() int {
	return m.Size()
}
func (m *Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Head proto.InternalMessageInfo

func (m *Head) GetConnID() string {
	if m != nil {
		return m.ConnID
	}
	return ""
}

func (m *Head) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *Head) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *Head) GetServerID() string {
	if m != nil {
		return m.ServerID
	}
	return ""
}

type BaseRes struct {
	ErrorCode int64 `protobuf:"varint,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
}

func (m *BaseRes) Reset()         { *m = BaseRes{} }
func (m *BaseRes) String() string { return proto.CompactTextString(m) }
func (*BaseRes) ProtoMessage()    {}
func (*BaseRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}
func (m *BaseRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRes.Merge(m, src)
}
func (m *BaseRes) XXX_Size() int {
	return m.Size()
}
func (m *BaseRes) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRes.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRes proto.InternalMessageInfo

func (m *BaseRes) GetErrorCode() int64 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type Msg struct {
	ApiId Api    `protobuf:"varint,1,opt,name=api_id,json=apiId,proto3,enum=base.Api" json:"api_id,omitempty"`
	Data  []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}
func (m *Msg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return m.Size()
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetApiId() Api {
	if m != nil {
		return m.ApiId
	}
	return Api_Empty
}

func (m *Msg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("base.Api", Api_name, Api_value)
	proto.RegisterType((*Head)(nil), "base.Head")
	proto.RegisterType((*BaseRes)(nil), "base.BaseRes")
	proto.RegisterType((*Msg)(nil), "base.Msg")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0x8d, 0xc9, 0x3f, 0x3c, 0x04, 0x58, 0x06, 0x08, 0xe6, 0x8f, 0xfc, 0x43, 0x1c, 0x7e, 0xad,
	0x38, 0x70, 0x68, 0x8f, 0x3d, 0x39, 0x31, 0x32, 0x51, 0xa9, 0x54, 0x4c, 0x51, 0x8f, 0x95, 0x61,
	0xa7, 0x74, 0xd5, 0xc4, 0xbb, 0x5a, 0x6f, 0x90, 0xfc, 0x2d, 0x7a, 0xef, 0x17, 0xea, 0x91, 0x63,
	0x8f, 0x15, 0x7c, 0x91, 0x6a, 0xd7, 0xa1, 0x38, 0x4a, 0xd5, 0x9b, 0x67, 0xde, 0xbc, 0xb7, 0xf3,
	0xde, 0xc8, 0x00, 0xd7, 0x59, 0x41, 0x27, 0x4a, 0x4b, 0x23, 0xb1, 0x65, 0xbf, 0x8f, 0xc6, 0xd0,
	0x3a, 0xa3, 0x8c, 0x63, 0x1f, 0x3a, 0x43, 0x99, 0xe7, 0xa3, 0x38, 0xf0, 0x0e, 0xbd, 0x97, 0x7e,
	0x3a, 0xab, 0x30, 0x80, 0x6e, 0xa2, 0xe5, 0x54, 0x8d, 0xe2, 0x60, 0xc9, 0x01, 0x4f, 0xa5, 0x65,
	0x44, 0x4a, 0xbd, 0xa5, 0x32, 0x68, 0x56, 0x8c, 0xaa, 0xc2, 0x3d, 0x58, 0xbe, 0x24, 0x7d, 0x47,
	0x7a, 0x14, 0x07, 0x2d, 0x87, 0xfc, 0xa9, 0x8f, 0x5e, 0x40, 0x77, 0x90, 0x15, 0x94, 0x52, 0x81,
	0x07, 0xe0, 0x9f, 0x6a, 0x2d, 0xf5, 0x50, 0x72, 0x72, 0x6f, 0x36, 0xd3, 0xe7, 0xc6, 0xd1, 0x1b,
	0x68, 0xbe, 0x2b, 0x6e, 0xf1, 0x10, 0x3a, 0x99, 0x12, 0x9f, 0x04, 0x77, 0x13, 0x6b, 0xaf, 0xfc,
	0x13, 0x67, 0x20, 0x52, 0x22, 0x6d, 0x67, 0x4a, 0x8c, 0x38, 0x22, 0xb4, 0xe2, 0xcc, 0x64, 0x6e,
	0xb9, 0x5e, 0xea, 0xbe, 0x8f, 0xbf, 0xb7, 0xa1, 0x19, 0x29, 0x81, 0x3e, 0xb4, 0x4f, 0x27, 0xca,
	0x94, 0xac, 0x81, 0xab, 0xe0, 0x0f, 0xa6, 0x62, 0xcc, 0xad, 0x2b, 0xe6, 0x61, 0x0f, 0x96, 0x3f,
	0xd2, 0xf5, 0xb9, 0xbc, 0x15, 0x39, 0x5b, 0xc2, 0x4d, 0x58, 0x8f, 0x38, 0x3f, 0x9d, 0xa8, 0xb1,
	0x2c, 0x89, 0x46, 0xf9, 0x67, 0xc9, 0x9a, 0xd8, 0x07, 0xbc, 0x52, 0x3c, 0x33, 0x34, 0xd7, 0x6f,
	0xd9, 0xe1, 0x84, 0xcc, 0x5c, 0xb3, 0x8d, 0xbb, 0xb0, 0x5d, 0x53, 0x38, 0xa3, 0x6c, 0x6c, 0xbe,
	0x38, 0xa8, 0x83, 0x07, 0x10, 0xcc, 0xeb, 0xd4, 0xd0, 0xae, 0x25, 0xd6, 0xd4, 0x6a, 0xd0, 0x32,
	0xee, 0xc3, 0xce, 0x82, 0x66, 0x4a, 0x37, 0x52, 0x73, 0xe6, 0x23, 0xc2, 0x5a, 0xc4, 0xf9, 0x50,
	0x4e, 0x54, 0x96, 0x97, 0x8e, 0x00, 0xb8, 0x0d, 0x1b, 0xd5, 0x4b, 0xf5, 0xf6, 0x8a, 0x1d, 0x4d,
	0xc8, 0xd4, 0x7b, 0x3d, 0xdc, 0x81, 0xcd, 0x88, 0xf3, 0xf7, 0x9a, 0xee, 0x28, 0x37, 0x42, 0xe6,
	0x33, 0xdd, 0x55, 0xdc, 0x83, 0x7e, 0xa5, 0xb1, 0x80, 0xad, 0x59, 0x52, 0x42, 0x66, 0x01, 0x58,
	0xc7, 0x0d, 0x58, 0x4d, 0xc8, 0x5c, 0x88, 0x5c, 0x4c, 0x3f, 0xc8, 0xaf, 0x94, 0x33, 0x86, 0x0c,
	0x7a, 0x6e, 0x56, 0xde, 0x89, 0xfc, 0x86, 0x0a, 0xb6, 0x81, 0x2b, 0xd0, 0xb5, 0x6b, 0x08, 0x53,
	0x32, 0xc4, 0x75, 0x58, 0x49, 0xc8, 0xc4, 0xa2, 0x30, 0x5a, 0xdc, 0x18, 0xb6, 0x69, 0xef, 0x93,
	0x90, 0xb9, 0x34, 0x9a, 0xc8, 0xb0, 0xad, 0x19, 0x7d, 0x28, 0x27, 0x93, 0x69, 0x6e, 0x19, 0xdb,
	0x36, 0xa8, 0x8b, 0x29, 0xe9, 0xb2, 0xca, 0xc3, 0xda, 0x18, 0x94, 0x57, 0x05, 0x69, 0xd6, 0xb7,
	0x97, 0x9a, 0x37, 0x78, 0x2e, 0x0a, 0xc3, 0x76, 0x6c, 0x1e, 0x31, 0x8d, 0x69, 0x3e, 0x8f, 0xc0,
	0xe6, 0xfa, 0x17, 0x1b, 0x8e, 0xb3, 0x8b, 0xff, 0xc1, 0xfe, 0xc2, 0x3d, 0x6a, 0x03, 0x7b, 0xb8,
	0x05, 0xac, 0x12, 0x7d, 0x5e, 0x84, 0xed, 0xe3, 0x31, 0xfc, 0xff, 0x0f, 0xda, 0xa0, 0x7c, 0xda,
	0x80, 0xb3, 0x83, 0x41, 0xf0, 0xe3, 0x21, 0xf4, 0xee, 0x1f, 0x42, 0xef, 0xd7, 0x43, 0xe8, 0x7d,
	0x7b, 0x0c, 0x1b, 0xf7, 0x8f, 0x61, 0xe3, 0xe7, 0x63, 0xd8, 0xb8, 0xee, 0xb8, 0x1f, 0xf3, 0xf5,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x7d, 0x30, 0x86, 0xa6, 0x03, 0x00, 0x00,
}

func (m *Head) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Head) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Head) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ServerID) > 0 {
		i -= len(m.ServerID)
		copy(dAtA[i:], m.ServerID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.ServerID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AppKey) > 0 {
		i -= len(m.AppKey)
		copy(dAtA[i:], m.AppKey)
		i = encodeVarintBase(dAtA, i, uint64(len(m.AppKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.GroupID) > 0 {
		i -= len(m.GroupID)
		copy(dAtA[i:], m.GroupID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.GroupID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnID) > 0 {
		i -= len(m.ConnID)
		copy(dAtA[i:], m.ConnID)
		i = encodeVarintBase(dAtA, i, uint64(len(m.ConnID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BaseRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ErrorCode != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.ErrorCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Msg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Msg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Msg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintBase(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.ApiId != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.ApiId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	offset -= sovBase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Head) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.GroupID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.AppKey)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.ServerID)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func (m *BaseRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrorCode != 0 {
		n += 1 + sovBase(uint64(m.ErrorCode))
	}
	return n
}

func (m *Msg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApiId != 0 {
		n += 1 + sovBase(uint64(m.ApiId))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func sovBase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Head) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: Head: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Head: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GroupID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
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
func (m *BaseRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: BaseRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
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
func (m *Msg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: Msg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Msg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiId", wireType)
			}
			m.ApiId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiId |= Api(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBase
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
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
				return 0, ErrInvalidLengthBase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBase = fmt.Errorf("proto: unexpected end of group")
)