// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: currency.proto

package currency

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_CAD Currencies = 1
	Currencies_HKD Currencies = 2
	Currencies_ISK Currencies = 3
	Currencies_PHP Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_HUF Currencies = 6
	Currencies_CZK Currencies = 7
	Currencies_AUD Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_IDR Currencies = 11
	Currencies_INR Currencies = 12
	Currencies_BRL Currencies = 13
	Currencies_RUB Currencies = 14
	Currencies_HRK Currencies = 15
	Currencies_JPY Currencies = 16
	Currencies_THB Currencies = 17
	Currencies_CHF Currencies = 18
	Currencies_SGD Currencies = 19
	Currencies_PLN Currencies = 20
	Currencies_BGN Currencies = 21
	Currencies_TRY Currencies = 22
	Currencies_CNY Currencies = 23
	Currencies_NOK Currencies = 24
	Currencies_NZD Currencies = 25
	Currencies_ZAR Currencies = 26
	Currencies_USD Currencies = 27
	Currencies_MXN Currencies = 28
	Currencies_ILS Currencies = 29
	Currencies_GBP Currencies = 30
	Currencies_KRW Currencies = 31
	Currencies_MYR Currencies = 32
)

// Enum value maps for Currencies.
var (
	Currencies_name = map[int32]string{
		0:  "EUR",
		1:  "CAD",
		2:  "HKD",
		3:  "ISK",
		4:  "PHP",
		5:  "DKK",
		6:  "HUF",
		7:  "CZK",
		8:  "AUD",
		9:  "RON",
		10: "SEK",
		11: "IDR",
		12: "INR",
		13: "BRL",
		14: "RUB",
		15: "HRK",
		16: "JPY",
		17: "THB",
		18: "CHF",
		19: "SGD",
		20: "PLN",
		21: "BGN",
		22: "TRY",
		23: "CNY",
		24: "NOK",
		25: "NZD",
		26: "ZAR",
		27: "USD",
		28: "MXN",
		29: "ILS",
		30: "GBP",
		31: "KRW",
		32: "MYR",
	}
	Currencies_value = map[string]int32{
		"EUR": 0,
		"CAD": 1,
		"HKD": 2,
		"ISK": 3,
		"PHP": 4,
		"DKK": 5,
		"HUF": 6,
		"CZK": 7,
		"AUD": 8,
		"RON": 9,
		"SEK": 10,
		"IDR": 11,
		"INR": 12,
		"BRL": 13,
		"RUB": 14,
		"HRK": 15,
		"JPY": 16,
		"THB": 17,
		"CHF": 18,
		"SGD": 19,
		"PLN": 20,
		"BGN": 21,
		"TRY": 22,
		"CNY": 23,
		"NOK": 24,
		"NZD": 25,
		"ZAR": 26,
		"USD": 27,
		"MXN": 28,
		"ILS": 29,
		"GBP": 30,
		"KRW": 31,
		"MYR": 32,
	}
)

func (x Currencies) Enum() *Currencies {
	p := new(Currencies)
	*p = x
	return p
}

func (x Currencies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currencies) Descriptor() protoreflect.EnumDescriptor {
	return file_currency_proto_enumTypes[0].Descriptor()
}

func (Currencies) Type() protoreflect.EnumType {
	return &file_currency_proto_enumTypes[0]
}

func (x Currencies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currencies.Descriptor instead.
func (Currencies) EnumDescriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

type RateRquest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=Currencies" json:"Destination,omitempty"`
}

func (x *RateRquest) Reset() {
	*x = RateRquest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRquest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRquest) ProtoMessage() {}

func (x *RateRquest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRquest.ProtoReflect.Descriptor instead.
func (*RateRquest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRquest) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_EUR
}

func (x *RateRquest) GetDestination() Currencies {
	if x != nil {
		return x.Destination
	}
	return Currencies_EUR
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate float64 `protobuf:"fixed64,1,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_currency_proto protoreflect.FileDescriptor

var file_currency_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5c, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x52, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x22,
	0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x52, 0x61,
	0x74, 0x65, 0x2a, 0xb5, 0x02, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x55, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x41,
	0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x4b, 0x44, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x49, 0x53, 0x4b, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x48, 0x50, 0x10, 0x04, 0x12, 0x07,
	0x0a, 0x03, 0x44, 0x4b, 0x4b, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x55, 0x46, 0x10, 0x06,
	0x12, 0x07, 0x0a, 0x03, 0x43, 0x5a, 0x4b, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x55, 0x44,
	0x10, 0x08, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x53,
	0x45, 0x4b, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x44, 0x52, 0x10, 0x0b, 0x12, 0x07, 0x0a,
	0x03, 0x49, 0x4e, 0x52, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x52, 0x4c, 0x10, 0x0d, 0x12,
	0x07, 0x0a, 0x03, 0x52, 0x55, 0x42, 0x10, 0x0e, 0x12, 0x07, 0x0a, 0x03, 0x48, 0x52, 0x4b, 0x10,
	0x0f, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x59, 0x10, 0x10, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x48,
	0x42, 0x10, 0x11, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x48, 0x46, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03,
	0x53, 0x47, 0x44, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4c, 0x4e, 0x10, 0x14, 0x12, 0x07,
	0x0a, 0x03, 0x42, 0x47, 0x4e, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x52, 0x59, 0x10, 0x16,
	0x12, 0x07, 0x0a, 0x03, 0x43, 0x4e, 0x59, 0x10, 0x17, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x4b,
	0x10, 0x18, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x5a, 0x44, 0x10, 0x19, 0x12, 0x07, 0x0a, 0x03, 0x5a,
	0x41, 0x52, 0x10, 0x1a, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x53, 0x44, 0x10, 0x1b, 0x12, 0x07, 0x0a,
	0x03, 0x4d, 0x58, 0x4e, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4c, 0x53, 0x10, 0x1d, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x42, 0x50, 0x10, 0x1e, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x52, 0x57, 0x10,
	0x1f, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x59, 0x52, 0x10, 0x20, 0x32, 0x31, 0x0a, 0x08, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x0b, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currency_proto_rawDescOnce sync.Once
	file_currency_proto_rawDescData = file_currency_proto_rawDesc
)

func file_currency_proto_rawDescGZIP() []byte {
	file_currency_proto_rawDescOnce.Do(func() {
		file_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_proto_rawDescData)
	})
	return file_currency_proto_rawDescData
}

var file_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_currency_proto_goTypes = []interface{}{
	(Currencies)(0),      // 0: Currencies
	(*RateRquest)(nil),   // 1: RateRquest
	(*RateResponse)(nil), // 2: RateResponse
}
var file_currency_proto_depIdxs = []int32{
	0, // 0: RateRquest.Base:type_name -> Currencies
	0, // 1: RateRquest.Destination:type_name -> Currencies
	1, // 2: Currency.GetRate:input_type -> RateRquest
	2, // 3: Currency.GetRate:output_type -> RateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_currency_proto_init() }
func file_currency_proto_init() {
	if File_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRquest); i {
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
		file_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
			RawDescriptor: file_currency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_proto_goTypes,
		DependencyIndexes: file_currency_proto_depIdxs,
		EnumInfos:         file_currency_proto_enumTypes,
		MessageInfos:      file_currency_proto_msgTypes,
	}.Build()
	File_currency_proto = out.File
	file_currency_proto_rawDesc = nil
	file_currency_proto_goTypes = nil
	file_currency_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrencyClient is the client API for Currency service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrencyClient interface {
	GetRate(ctx context.Context, in *RateRquest, opts ...grpc.CallOption) (*RateResponse, error)
}

type currencyClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyClient(cc grpc.ClientConnInterface) CurrencyClient {
	return &currencyClient{cc}
}

func (c *currencyClient) GetRate(ctx context.Context, in *RateRquest, opts ...grpc.CallOption) (*RateResponse, error) {
	out := new(RateResponse)
	err := c.cc.Invoke(ctx, "/Currency/GetRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServer is the server API for Currency service.
type CurrencyServer interface {
	GetRate(context.Context, *RateRquest) (*RateResponse, error)
}

// UnimplementedCurrencyServer can be embedded to have forward compatible implementations.
type UnimplementedCurrencyServer struct {
}

func (*UnimplementedCurrencyServer) GetRate(context.Context, *RateRquest) (*RateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRate not implemented")
}

func RegisterCurrencyServer(s *grpc.Server, srv CurrencyServer) {
	s.RegisterService(&_Currency_serviceDesc, srv)
}

func _Currency_GetRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRquest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServer).GetRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Currency/GetRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServer).GetRate(ctx, req.(*RateRquest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Currency_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Currency",
	HandlerType: (*CurrencyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRate",
			Handler:    _Currency_GetRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency.proto",
}
