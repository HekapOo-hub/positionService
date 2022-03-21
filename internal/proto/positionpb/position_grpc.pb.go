// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package positionpb

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

// PositionServiceClient is the client API for PositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PositionServiceClient interface {
	Close(ctx context.Context, in *PositionID, opts ...grpc.CallOption) (*Empty, error)
	Open(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Empty, error)
	GetOpen(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*Positions, error)
	UpdatePrices(ctx context.Context, opts ...grpc.CallOption) (PositionService_UpdatePricesClient, error)
	GetProfitLoss(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (PositionService_GetProfitLossClient, error)
}

type positionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPositionServiceClient(cc grpc.ClientConnInterface) PositionServiceClient {
	return &positionServiceClient{cc}
}

func (c *positionServiceClient) Close(ctx context.Context, in *PositionID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/position_proto.PositionService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) Open(ctx context.Context, in *Position, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/position_proto.PositionService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) GetOpen(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (*Positions, error) {
	out := new(Positions)
	err := c.cc.Invoke(ctx, "/position_proto.PositionService/GetOpen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *positionServiceClient) UpdatePrices(ctx context.Context, opts ...grpc.CallOption) (PositionService_UpdatePricesClient, error) {
	stream, err := c.cc.NewStream(ctx, &PositionService_ServiceDesc.Streams[0], "/position_proto.PositionService/UpdatePrices", opts...)
	if err != nil {
		return nil, err
	}
	x := &positionServiceUpdatePricesClient{stream}
	return x, nil
}

type PositionService_UpdatePricesClient interface {
	Send(*Price) error
	Recv() (*Empty, error)
	grpc.ClientStream
}

type positionServiceUpdatePricesClient struct {
	grpc.ClientStream
}

func (x *positionServiceUpdatePricesClient) Send(m *Price) error {
	return x.ClientStream.SendMsg(m)
}

func (x *positionServiceUpdatePricesClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *positionServiceClient) GetProfitLoss(ctx context.Context, in *AccountID, opts ...grpc.CallOption) (PositionService_GetProfitLossClient, error) {
	stream, err := c.cc.NewStream(ctx, &PositionService_ServiceDesc.Streams[1], "/position_proto.PositionService/GetProfitLoss", opts...)
	if err != nil {
		return nil, err
	}
	x := &positionServiceGetProfitLossClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PositionService_GetProfitLossClient interface {
	Recv() (*ProfitLoss, error)
	grpc.ClientStream
}

type positionServiceGetProfitLossClient struct {
	grpc.ClientStream
}

func (x *positionServiceGetProfitLossClient) Recv() (*ProfitLoss, error) {
	m := new(ProfitLoss)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PositionServiceServer is the server API for PositionService service.
// All implementations must embed UnimplementedPositionServiceServer
// for forward compatibility
type PositionServiceServer interface {
	Close(context.Context, *PositionID) (*Empty, error)
	Open(context.Context, *Position) (*Empty, error)
	GetOpen(context.Context, *AccountID) (*Positions, error)
	UpdatePrices(PositionService_UpdatePricesServer) error
	GetProfitLoss(*AccountID, PositionService_GetProfitLossServer) error
	mustEmbedUnimplementedPositionServiceServer()
}

// UnimplementedPositionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPositionServiceServer struct {
}

func (UnimplementedPositionServiceServer) Close(context.Context, *PositionID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedPositionServiceServer) Open(context.Context, *Position) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedPositionServiceServer) GetOpen(context.Context, *AccountID) (*Positions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOpen not implemented")
}
func (UnimplementedPositionServiceServer) UpdatePrices(PositionService_UpdatePricesServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdatePrices not implemented")
}
func (UnimplementedPositionServiceServer) GetProfitLoss(*AccountID, PositionService_GetProfitLossServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProfitLoss not implemented")
}
func (UnimplementedPositionServiceServer) mustEmbedUnimplementedPositionServiceServer() {}

// UnsafePositionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PositionServiceServer will
// result in compilation errors.
type UnsafePositionServiceServer interface {
	mustEmbedUnimplementedPositionServiceServer()
}

func RegisterPositionServiceServer(s grpc.ServiceRegistrar, srv PositionServiceServer) {
	s.RegisterService(&PositionService_ServiceDesc, srv)
}

func _PositionService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_proto.PositionService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Close(ctx, req.(*PositionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Position)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_proto.PositionService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).Open(ctx, req.(*Position))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_GetOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PositionServiceServer).GetOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/position_proto.PositionService/GetOpen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PositionServiceServer).GetOpen(ctx, req.(*AccountID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PositionService_UpdatePrices_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PositionServiceServer).UpdatePrices(&positionServiceUpdatePricesServer{stream})
}

type PositionService_UpdatePricesServer interface {
	Send(*Empty) error
	Recv() (*Price, error)
	grpc.ServerStream
}

type positionServiceUpdatePricesServer struct {
	grpc.ServerStream
}

func (x *positionServiceUpdatePricesServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *positionServiceUpdatePricesServer) Recv() (*Price, error) {
	m := new(Price)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PositionService_GetProfitLoss_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AccountID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PositionServiceServer).GetProfitLoss(m, &positionServiceGetProfitLossServer{stream})
}

type PositionService_GetProfitLossServer interface {
	Send(*ProfitLoss) error
	grpc.ServerStream
}

type positionServiceGetProfitLossServer struct {
	grpc.ServerStream
}

func (x *positionServiceGetProfitLossServer) Send(m *ProfitLoss) error {
	return x.ServerStream.SendMsg(m)
}

// PositionService_ServiceDesc is the grpc.ServiceDesc for PositionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PositionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "position_proto.PositionService",
	HandlerType: (*PositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Close",
			Handler:    _PositionService_Close_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _PositionService_Open_Handler,
		},
		{
			MethodName: "GetOpen",
			Handler:    _PositionService_GetOpen_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdatePrices",
			Handler:       _PositionService_UpdatePrices_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetProfitLoss",
			Handler:       _PositionService_GetProfitLoss_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "position.proto",
}