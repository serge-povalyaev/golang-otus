// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: internal/server/grpc/proto/event.proto

package grpc

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

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Error, error)
	RemoveEvent(ctx context.Context, in *RemoveEventRequest, opts ...grpc.CallOption) (*Error, error)
	GetUserEvents(ctx context.Context, in *GetUserEventsRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error)
	GetUserEventsByDay(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error)
	GetUserEventsByWeek(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error)
	GetUserEventsByMonth(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.EventService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) RemoveEvent(ctx context.Context, in *RemoveEventRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/api.EventService/RemoveEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetUserEvents(ctx context.Context, in *GetUserEventsRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error) {
	out := new(GetUserEventsResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetUserEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetUserEventsByDay(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error) {
	out := new(GetUserEventsResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetUserEventsByDay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetUserEventsByWeek(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error) {
	out := new(GetUserEventsResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetUserEventsByWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) GetUserEventsByMonth(ctx context.Context, in *GetUserEventsShortRequest, opts ...grpc.CallOption) (*GetUserEventsResponse, error) {
	out := new(GetUserEventsResponse)
	err := c.cc.Invoke(ctx, "/api.EventService/GetUserEventsByMonth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*Error, error)
	RemoveEvent(context.Context, *RemoveEventRequest) (*Error, error)
	GetUserEvents(context.Context, *GetUserEventsRequest) (*GetUserEventsResponse, error)
	GetUserEventsByDay(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error)
	GetUserEventsByWeek(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error)
	GetUserEventsByMonth(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedEventServiceServer) GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventServiceServer) UpdateEvent(context.Context, *UpdateEventRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (UnimplementedEventServiceServer) RemoveEvent(context.Context, *RemoveEventRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEvent not implemented")
}
func (UnimplementedEventServiceServer) GetUserEvents(context.Context, *GetUserEventsRequest) (*GetUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEvents not implemented")
}
func (UnimplementedEventServiceServer) GetUserEventsByDay(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEventsByDay not implemented")
}
func (UnimplementedEventServiceServer) GetUserEventsByWeek(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEventsByWeek not implemented")
}
func (UnimplementedEventServiceServer) GetUserEventsByMonth(context.Context, *GetUserEventsShortRequest) (*GetUserEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEventsByMonth not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_RemoveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).RemoveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/RemoveEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).RemoveEvent(ctx, req.(*RemoveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetUserEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetUserEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetUserEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetUserEvents(ctx, req.(*GetUserEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetUserEventsByDay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEventsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetUserEventsByDay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetUserEventsByDay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetUserEventsByDay(ctx, req.(*GetUserEventsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetUserEventsByWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEventsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetUserEventsByWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetUserEventsByWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetUserEventsByWeek(ctx, req.(*GetUserEventsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_GetUserEventsByMonth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEventsShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetUserEventsByMonth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.EventService/GetUserEventsByMonth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetUserEventsByMonth(ctx, req.(*GetUserEventsShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _EventService_AddEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _EventService_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventService_UpdateEvent_Handler,
		},
		{
			MethodName: "RemoveEvent",
			Handler:    _EventService_RemoveEvent_Handler,
		},
		{
			MethodName: "GetUserEvents",
			Handler:    _EventService_GetUserEvents_Handler,
		},
		{
			MethodName: "GetUserEventsByDay",
			Handler:    _EventService_GetUserEventsByDay_Handler,
		},
		{
			MethodName: "GetUserEventsByWeek",
			Handler:    _EventService_GetUserEventsByWeek_Handler,
		},
		{
			MethodName: "GetUserEventsByMonth",
			Handler:    _EventService_GetUserEventsByMonth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/server/grpc/proto/event.proto",
}