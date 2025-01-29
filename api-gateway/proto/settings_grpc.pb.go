// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: settings.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SettingsService_GetBotSettings_FullMethodName      = "/settings.SettingsService/GetBotSettings"
	SettingsService_UpdateBotSettings_FullMethodName   = "/settings.SettingsService/UpdateBotSettings"
	SettingsService_GetSettingsByGuild_FullMethodName  = "/settings.SettingsService/GetSettingsByGuild"
	SettingsService_GetAllGuildSettings_FullMethodName = "/settings.SettingsService/GetAllGuildSettings"
	SettingsService_UpdateGuildSettings_FullMethodName = "/settings.SettingsService/UpdateGuildSettings"
	SettingsService_DeleteGuildSetting_FullMethodName  = "/settings.SettingsService/DeleteGuildSetting"
)

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис SettingsService с методами для работы с настройками
type SettingsServiceClient interface {
	// Методы для работы с настройками бота
	GetBotSettings(ctx context.Context, in *GetBotSettingsRequest, opts ...grpc.CallOption) (*GetBotSettingsResponse, error)
	UpdateBotSettings(ctx context.Context, in *UpdateBotSettingsRequest, opts ...grpc.CallOption) (*UpdateBotSettingsResponse, error)
	// Методы для работы с настройками гильдии
	GetSettingsByGuild(ctx context.Context, in *GetSettingsByGuildRequest, opts ...grpc.CallOption) (*GetSettingsByGuildResponse, error)
	GetAllGuildSettings(ctx context.Context, in *GetAllGuildSettingsRequest, opts ...grpc.CallOption) (*GetAllGuildSettingsResponse, error)
	UpdateGuildSettings(ctx context.Context, in *UpdateGuildSettingsRequest, opts ...grpc.CallOption) (*UpdateGuildSettingsResponse, error)
	DeleteGuildSetting(ctx context.Context, in *DeleteGuildSettingRequest, opts ...grpc.CallOption) (*DeleteGuildSettingResponse, error)
}

type settingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) GetBotSettings(ctx context.Context, in *GetBotSettingsRequest, opts ...grpc.CallOption) (*GetBotSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBotSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetBotSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) UpdateBotSettings(ctx context.Context, in *UpdateBotSettingsRequest, opts ...grpc.CallOption) (*UpdateBotSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBotSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_UpdateBotSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetSettingsByGuild(ctx context.Context, in *GetSettingsByGuildRequest, opts ...grpc.CallOption) (*GetSettingsByGuildResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSettingsByGuildResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetSettingsByGuild_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) GetAllGuildSettings(ctx context.Context, in *GetAllGuildSettingsRequest, opts ...grpc.CallOption) (*GetAllGuildSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllGuildSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_GetAllGuildSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) UpdateGuildSettings(ctx context.Context, in *UpdateGuildSettingsRequest, opts ...grpc.CallOption) (*UpdateGuildSettingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateGuildSettingsResponse)
	err := c.cc.Invoke(ctx, SettingsService_UpdateGuildSettings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) DeleteGuildSetting(ctx context.Context, in *DeleteGuildSettingRequest, opts ...grpc.CallOption) (*DeleteGuildSettingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteGuildSettingResponse)
	err := c.cc.Invoke(ctx, SettingsService_DeleteGuildSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
// All implementations must embed UnimplementedSettingsServiceServer
// for forward compatibility.
//
// Сервис SettingsService с методами для работы с настройками
type SettingsServiceServer interface {
	// Методы для работы с настройками бота
	GetBotSettings(context.Context, *GetBotSettingsRequest) (*GetBotSettingsResponse, error)
	UpdateBotSettings(context.Context, *UpdateBotSettingsRequest) (*UpdateBotSettingsResponse, error)
	// Методы для работы с настройками гильдии
	GetSettingsByGuild(context.Context, *GetSettingsByGuildRequest) (*GetSettingsByGuildResponse, error)
	GetAllGuildSettings(context.Context, *GetAllGuildSettingsRequest) (*GetAllGuildSettingsResponse, error)
	UpdateGuildSettings(context.Context, *UpdateGuildSettingsRequest) (*UpdateGuildSettingsResponse, error)
	DeleteGuildSetting(context.Context, *DeleteGuildSettingRequest) (*DeleteGuildSettingResponse, error)
	mustEmbedUnimplementedSettingsServiceServer()
}

// UnimplementedSettingsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSettingsServiceServer struct{}

func (UnimplementedSettingsServiceServer) GetBotSettings(context.Context, *GetBotSettingsRequest) (*GetBotSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBotSettings not implemented")
}
func (UnimplementedSettingsServiceServer) UpdateBotSettings(context.Context, *UpdateBotSettingsRequest) (*UpdateBotSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBotSettings not implemented")
}
func (UnimplementedSettingsServiceServer) GetSettingsByGuild(context.Context, *GetSettingsByGuildRequest) (*GetSettingsByGuildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettingsByGuild not implemented")
}
func (UnimplementedSettingsServiceServer) GetAllGuildSettings(context.Context, *GetAllGuildSettingsRequest) (*GetAllGuildSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGuildSettings not implemented")
}
func (UnimplementedSettingsServiceServer) UpdateGuildSettings(context.Context, *UpdateGuildSettingsRequest) (*UpdateGuildSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuildSettings not implemented")
}
func (UnimplementedSettingsServiceServer) DeleteGuildSetting(context.Context, *DeleteGuildSettingRequest) (*DeleteGuildSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGuildSetting not implemented")
}
func (UnimplementedSettingsServiceServer) mustEmbedUnimplementedSettingsServiceServer() {}
func (UnimplementedSettingsServiceServer) testEmbeddedByValue()                         {}

// UnsafeSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServiceServer will
// result in compilation errors.
type UnsafeSettingsServiceServer interface {
	mustEmbedUnimplementedSettingsServiceServer()
}

func RegisterSettingsServiceServer(s grpc.ServiceRegistrar, srv SettingsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSettingsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SettingsService_ServiceDesc, srv)
}

func _SettingsService_GetBotSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBotSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetBotSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetBotSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetBotSettings(ctx, req.(*GetBotSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_UpdateBotSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBotSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).UpdateBotSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_UpdateBotSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).UpdateBotSettings(ctx, req.(*UpdateBotSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetSettingsByGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsByGuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetSettingsByGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetSettingsByGuild_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetSettingsByGuild(ctx, req.(*GetSettingsByGuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_GetAllGuildSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGuildSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).GetAllGuildSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_GetAllGuildSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).GetAllGuildSettings(ctx, req.(*GetAllGuildSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_UpdateGuildSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuildSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).UpdateGuildSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_UpdateGuildSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).UpdateGuildSettings(ctx, req.(*UpdateGuildSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_DeleteGuildSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGuildSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).DeleteGuildSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_DeleteGuildSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).DeleteGuildSetting(ctx, req.(*DeleteGuildSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingsService_ServiceDesc is the grpc.ServiceDesc for SettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "settings.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBotSettings",
			Handler:    _SettingsService_GetBotSettings_Handler,
		},
		{
			MethodName: "UpdateBotSettings",
			Handler:    _SettingsService_UpdateBotSettings_Handler,
		},
		{
			MethodName: "GetSettingsByGuild",
			Handler:    _SettingsService_GetSettingsByGuild_Handler,
		},
		{
			MethodName: "GetAllGuildSettings",
			Handler:    _SettingsService_GetAllGuildSettings_Handler,
		},
		{
			MethodName: "UpdateGuildSettings",
			Handler:    _SettingsService_UpdateGuildSettings_Handler,
		},
		{
			MethodName: "DeleteGuildSetting",
			Handler:    _SettingsService_DeleteGuildSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "settings.proto",
}
