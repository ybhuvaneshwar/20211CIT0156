

package calculator

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

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=FirstNumber,proto3" json:"FirstNumber,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=SecondNumber,proto3" json:"SecondNumber,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  float32 `protobuf:"fixed32,1,opt,name=FirstNumber,proto3" json:"FirstNumber,omitempty"`
	SecondNumber float32 `protobuf:"fixed32,2,opt,name=SecondNumber,proto3" json:"SecondNumber,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *AverageRequest) GetFirstNumber() float32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *AverageRequest) GetSecondNumber() float32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *AverageResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x56, 0x0a,
	0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x66, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x25,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x12, 0x0b, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculator_proto_goTypes = []interface{}{
	(*SumRequest)(nil),      // 0: SumRequest
	(*SumResponse)(nil),     // 1: SumResponse
	(*AverageRequest)(nil),  // 2: AverageRequest
	(*AverageResponse)(nil), // 3: AverageResponse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: Calculator.GetSum:input_type -> SumRequest
	2, // 1: Calculator.GetAverage:input_type -> AverageRequest
	1, // 2: Calculator.GetSum:output_type -> SumResponse
	3, // 3: Calculator.GetAverage:output_type -> AverageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
		file_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	GetSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	GetAverage(ctx context.Context, in *AverageRequest, opts ...grpc.CallOption) (*AverageResponse, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) GetSum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/Calculator/GetSum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorClient) GetAverage(ctx context.Context, in *AverageRequest, opts ...grpc.CallOption) (*AverageResponse, error) {
	out := new(AverageResponse)
	err := c.cc.Invoke(ctx, "/Calculator/GetAverage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	GetSum(context.Context, *SumRequest) (*SumResponse, error)
	GetAverage(context.Context, *AverageRequest) (*AverageResponse, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) GetSum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSum not implemented")
}
func (*UnimplementedCalculatorServer) GetAverage(context.Context, *AverageRequest) (*AverageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAverage not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_GetSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).GetSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/GetSum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).GetSum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculator_GetAverage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AverageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).GetAverage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/GetAverage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).GetAverage(ctx, req.(*AverageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSum",
			Handler:    _Calculator_GetSum_Handler,
		},
		{
			MethodName: "GetAverage",
			Handler:    _Calculator_GetAverage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
