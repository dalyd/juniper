// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: system_metrics.proto

package gopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CedarSystemMetricsClient is the client API for CedarSystemMetrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CedarSystemMetricsClient interface {
	CreateSystemMetricsRecord(ctx context.Context, in *SystemMetrics, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
	AddSystemMetrics(ctx context.Context, in *SystemMetricsData, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
	StreamSystemMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarSystemMetrics_StreamSystemMetricsClient, error)
	CloseMetrics(ctx context.Context, in *SystemMetricsSeriesEnd, opts ...grpc.CallOption) (*SystemMetricsResponse, error)
}

type cedarSystemMetricsClient struct {
	cc grpc.ClientConnInterface
}

func NewCedarSystemMetricsClient(cc grpc.ClientConnInterface) CedarSystemMetricsClient {
	return &cedarSystemMetricsClient{cc}
}

func (c *cedarSystemMetricsClient) CreateSystemMetricsRecord(ctx context.Context, in *SystemMetrics, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/CreateSystemMetricsRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarSystemMetricsClient) AddSystemMetrics(ctx context.Context, in *SystemMetricsData, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/AddSystemMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cedarSystemMetricsClient) StreamSystemMetrics(ctx context.Context, opts ...grpc.CallOption) (CedarSystemMetrics_StreamSystemMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CedarSystemMetrics_ServiceDesc.Streams[0], "/cedar.CedarSystemMetrics/StreamSystemMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &cedarSystemMetricsStreamSystemMetricsClient{stream}
	return x, nil
}

type CedarSystemMetrics_StreamSystemMetricsClient interface {
	Send(*SystemMetricsData) error
	CloseAndRecv() (*SystemMetricsResponse, error)
	grpc.ClientStream
}

type cedarSystemMetricsStreamSystemMetricsClient struct {
	grpc.ClientStream
}

func (x *cedarSystemMetricsStreamSystemMetricsClient) Send(m *SystemMetricsData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cedarSystemMetricsStreamSystemMetricsClient) CloseAndRecv() (*SystemMetricsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SystemMetricsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cedarSystemMetricsClient) CloseMetrics(ctx context.Context, in *SystemMetricsSeriesEnd, opts ...grpc.CallOption) (*SystemMetricsResponse, error) {
	out := new(SystemMetricsResponse)
	err := c.cc.Invoke(ctx, "/cedar.CedarSystemMetrics/CloseMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CedarSystemMetricsServer is the server API for CedarSystemMetrics service.
// All implementations must embed UnimplementedCedarSystemMetricsServer
// for forward compatibility
type CedarSystemMetricsServer interface {
	CreateSystemMetricsRecord(context.Context, *SystemMetrics) (*SystemMetricsResponse, error)
	AddSystemMetrics(context.Context, *SystemMetricsData) (*SystemMetricsResponse, error)
	StreamSystemMetrics(CedarSystemMetrics_StreamSystemMetricsServer) error
	CloseMetrics(context.Context, *SystemMetricsSeriesEnd) (*SystemMetricsResponse, error)
	mustEmbedUnimplementedCedarSystemMetricsServer()
}

// UnimplementedCedarSystemMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedCedarSystemMetricsServer struct {
}

func (UnimplementedCedarSystemMetricsServer) CreateSystemMetricsRecord(context.Context, *SystemMetrics) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSystemMetricsRecord not implemented")
}
func (UnimplementedCedarSystemMetricsServer) AddSystemMetrics(context.Context, *SystemMetricsData) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSystemMetrics not implemented")
}
func (UnimplementedCedarSystemMetricsServer) StreamSystemMetrics(CedarSystemMetrics_StreamSystemMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSystemMetrics not implemented")
}
func (UnimplementedCedarSystemMetricsServer) CloseMetrics(context.Context, *SystemMetricsSeriesEnd) (*SystemMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseMetrics not implemented")
}
func (UnimplementedCedarSystemMetricsServer) mustEmbedUnimplementedCedarSystemMetricsServer() {}

// UnsafeCedarSystemMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CedarSystemMetricsServer will
// result in compilation errors.
type UnsafeCedarSystemMetricsServer interface {
	mustEmbedUnimplementedCedarSystemMetricsServer()
}

func RegisterCedarSystemMetricsServer(s grpc.ServiceRegistrar, srv CedarSystemMetricsServer) {
	s.RegisterService(&CedarSystemMetrics_ServiceDesc, srv)
}

func _CedarSystemMetrics_CreateSystemMetricsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).CreateSystemMetricsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/CreateSystemMetricsRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).CreateSystemMetricsRecord(ctx, req.(*SystemMetrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarSystemMetrics_AddSystemMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetricsData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).AddSystemMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/AddSystemMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).AddSystemMetrics(ctx, req.(*SystemMetricsData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CedarSystemMetrics_StreamSystemMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CedarSystemMetricsServer).StreamSystemMetrics(&cedarSystemMetricsStreamSystemMetricsServer{stream})
}

type CedarSystemMetrics_StreamSystemMetricsServer interface {
	SendAndClose(*SystemMetricsResponse) error
	Recv() (*SystemMetricsData, error)
	grpc.ServerStream
}

type cedarSystemMetricsStreamSystemMetricsServer struct {
	grpc.ServerStream
}

func (x *cedarSystemMetricsStreamSystemMetricsServer) SendAndClose(m *SystemMetricsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cedarSystemMetricsStreamSystemMetricsServer) Recv() (*SystemMetricsData, error) {
	m := new(SystemMetricsData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CedarSystemMetrics_CloseMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemMetricsSeriesEnd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CedarSystemMetricsServer).CloseMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cedar.CedarSystemMetrics/CloseMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CedarSystemMetricsServer).CloseMetrics(ctx, req.(*SystemMetricsSeriesEnd))
	}
	return interceptor(ctx, in, info, handler)
}

// CedarSystemMetrics_ServiceDesc is the grpc.ServiceDesc for CedarSystemMetrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CedarSystemMetrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cedar.CedarSystemMetrics",
	HandlerType: (*CedarSystemMetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSystemMetricsRecord",
			Handler:    _CedarSystemMetrics_CreateSystemMetricsRecord_Handler,
		},
		{
			MethodName: "AddSystemMetrics",
			Handler:    _CedarSystemMetrics_AddSystemMetrics_Handler,
		},
		{
			MethodName: "CloseMetrics",
			Handler:    _CedarSystemMetrics_CloseMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSystemMetrics",
			Handler:       _CedarSystemMetrics_StreamSystemMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "system_metrics.proto",
}
