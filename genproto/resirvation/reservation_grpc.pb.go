// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: reservation.proto

package resirvation

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

// ResirvationClient is the client API for Resirvation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResirvationClient interface {
	Createreservations(ctx context.Context, in *RequestReservations, opts ...grpc.CallOption) (*Status, error)
	GetAllReservations(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Reservations, error)
	GetByIdReservations(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*Reservation, error)
	UpdateReservations(ctx context.Context, in *ReservationUpdate, opts ...grpc.CallOption) (*Status, error)
	DeleteReservations(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*Status, error)
	GetReservationsByUserId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Reservations, error)
	OrderMeal(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Status, error)
	PayForReservation(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*Status, error)
}

type resirvationClient struct {
	cc grpc.ClientConnInterface
}

func NewResirvationClient(cc grpc.ClientConnInterface) ResirvationClient {
	return &resirvationClient{cc}
}

func (c *resirvationClient) Createreservations(ctx context.Context, in *RequestReservations, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/Createreservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) GetAllReservations(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*Reservations, error) {
	out := new(Reservations)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/GetAllReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) GetByIdReservations(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*Reservation, error) {
	out := new(Reservation)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/GetByIdReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) UpdateReservations(ctx context.Context, in *ReservationUpdate, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/UpdateReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) DeleteReservations(ctx context.Context, in *ReservationId, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/DeleteReservations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) GetReservationsByUserId(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Reservations, error) {
	out := new(Reservations)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/GetReservationsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) OrderMeal(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/OrderMeal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resirvationClient) PayForReservation(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/resirvation.Resirvation/PayForReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResirvationServer is the server API for Resirvation service.
// All implementations must embed UnimplementedResirvationServer
// for forward compatibility
type ResirvationServer interface {
	Createreservations(context.Context, *RequestReservations) (*Status, error)
	GetAllReservations(context.Context, *Filter) (*Reservations, error)
	GetByIdReservations(context.Context, *ReservationId) (*Reservation, error)
	UpdateReservations(context.Context, *ReservationUpdate) (*Status, error)
	DeleteReservations(context.Context, *ReservationId) (*Status, error)
	GetReservationsByUserId(context.Context, *UserId) (*Reservations, error)
	OrderMeal(context.Context, *Order) (*Status, error)
	PayForReservation(context.Context, *Payment) (*Status, error)
	mustEmbedUnimplementedResirvationServer()
}

// UnimplementedResirvationServer must be embedded to have forward compatible implementations.
type UnimplementedResirvationServer struct {
}

func (UnimplementedResirvationServer) Createreservations(context.Context, *RequestReservations) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Createreservations not implemented")
}
func (UnimplementedResirvationServer) GetAllReservations(context.Context, *Filter) (*Reservations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedResirvationServer) GetByIdReservations(context.Context, *ReservationId) (*Reservation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdReservations not implemented")
}
func (UnimplementedResirvationServer) UpdateReservations(context.Context, *ReservationUpdate) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservations not implemented")
}
func (UnimplementedResirvationServer) DeleteReservations(context.Context, *ReservationId) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservations not implemented")
}
func (UnimplementedResirvationServer) GetReservationsByUserId(context.Context, *UserId) (*Reservations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservationsByUserId not implemented")
}
func (UnimplementedResirvationServer) OrderMeal(context.Context, *Order) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderMeal not implemented")
}
func (UnimplementedResirvationServer) PayForReservation(context.Context, *Payment) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayForReservation not implemented")
}
func (UnimplementedResirvationServer) mustEmbedUnimplementedResirvationServer() {}

// UnsafeResirvationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResirvationServer will
// result in compilation errors.
type UnsafeResirvationServer interface {
	mustEmbedUnimplementedResirvationServer()
}

func RegisterResirvationServer(s grpc.ServiceRegistrar, srv ResirvationServer) {
	s.RegisterService(&Resirvation_ServiceDesc, srv)
}

func _Resirvation_Createreservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestReservations)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).Createreservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/Createreservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).Createreservations(ctx, req.(*RequestReservations))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/GetAllReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).GetAllReservations(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_GetByIdReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).GetByIdReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/GetByIdReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).GetByIdReservations(ctx, req.(*ReservationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_UpdateReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).UpdateReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/UpdateReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).UpdateReservations(ctx, req.(*ReservationUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_DeleteReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).DeleteReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/DeleteReservations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).DeleteReservations(ctx, req.(*ReservationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_GetReservationsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).GetReservationsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/GetReservationsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).GetReservationsByUserId(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_OrderMeal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).OrderMeal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/OrderMeal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).OrderMeal(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resirvation_PayForReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResirvationServer).PayForReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resirvation.Resirvation/PayForReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResirvationServer).PayForReservation(ctx, req.(*Payment))
	}
	return interceptor(ctx, in, info, handler)
}

// Resirvation_ServiceDesc is the grpc.ServiceDesc for Resirvation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resirvation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resirvation.Resirvation",
	HandlerType: (*ResirvationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Createreservations",
			Handler:    _Resirvation_Createreservations_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _Resirvation_GetAllReservations_Handler,
		},
		{
			MethodName: "GetByIdReservations",
			Handler:    _Resirvation_GetByIdReservations_Handler,
		},
		{
			MethodName: "UpdateReservations",
			Handler:    _Resirvation_UpdateReservations_Handler,
		},
		{
			MethodName: "DeleteReservations",
			Handler:    _Resirvation_DeleteReservations_Handler,
		},
		{
			MethodName: "GetReservationsByUserId",
			Handler:    _Resirvation_GetReservationsByUserId_Handler,
		},
		{
			MethodName: "OrderMeal",
			Handler:    _Resirvation_OrderMeal_Handler,
		},
		{
			MethodName: "PayForReservation",
			Handler:    _Resirvation_PayForReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation.proto",
}
