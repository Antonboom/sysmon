// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MonitorClient is the client API for Monitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitorClient interface {
	GetStats(ctx context.Context, in *MonRequest, opts ...grpc.CallOption) (Monitor_GetStatsClient, error)
}

type monitorClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorClient(cc grpc.ClientConnInterface) MonitorClient {
	return &monitorClient{cc}
}

func (c *monitorClient) GetStats(ctx context.Context, in *MonRequest, opts ...grpc.CallOption) (Monitor_GetStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Monitor_ServiceDesc.Streams[0], "/monitor.Monitor/GetStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorGetStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Monitor_GetStatsClient interface {
	Recv() (*StatSnapshot, error)
	grpc.ClientStream
}

type monitorGetStatsClient struct {
	grpc.ClientStream
}

func (x *monitorGetStatsClient) Recv() (*StatSnapshot, error) {
	m := new(StatSnapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorServer is the server API for Monitor service.
// All implementations must embed UnimplementedMonitorServer
// for forward compatibility
type MonitorServer interface {
	GetStats(*MonRequest, Monitor_GetStatsServer) error
	mustEmbedUnimplementedMonitorServer()
}

// UnimplementedMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedMonitorServer struct {
}

func (UnimplementedMonitorServer) GetStats(*MonRequest, Monitor_GetStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedMonitorServer) mustEmbedUnimplementedMonitorServer() {}

// UnsafeMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitorServer will
// result in compilation errors.
type UnsafeMonitorServer interface {
	mustEmbedUnimplementedMonitorServer()
}

func RegisterMonitorServer(s grpc.ServiceRegistrar, srv MonitorServer) {
	s.RegisterService(&Monitor_ServiceDesc, srv)
}

func _Monitor_GetStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorServer).GetStats(m, &monitorGetStatsServer{stream})
}

type Monitor_GetStatsServer interface {
	Send(*StatSnapshot) error
	grpc.ServerStream
}

type monitorGetStatsServer struct {
	grpc.ServerStream
}

func (x *monitorGetStatsServer) Send(m *StatSnapshot) error {
	return x.ServerStream.SendMsg(m)
}

// Monitor_ServiceDesc is the grpc.ServiceDesc for Monitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitor.Monitor",
	HandlerType: (*MonitorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStats",
			Handler:       _Monitor_GetStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sysmon.proto",
}
