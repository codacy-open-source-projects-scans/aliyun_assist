// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: commander.proto

package agrpc

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

// CommanderClient is the client API for Commander service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommanderClient interface {
	// For handshake
	Handshake(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*HandshakeResp, error)
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
	ConnectFromAgent(ctx context.Context, in *ConnectFromAgentReq, opts ...grpc.CallOption) (*ConnectFromAgentResp, error)
	// For execute task
	PreCheckScript(ctx context.Context, in *PreCheckScriptReq, opts ...grpc.CallOption) (*PreCheckScriptResp, error)
	PrepareScript(ctx context.Context, in *PrepareScriptReq, opts ...grpc.CallOption) (*PrepareScriptResp, error)
	SubmitScript(ctx context.Context, in *SubmitScriptReq, opts ...grpc.CallOption) (*SubmitScriptResp, error)
	CancelSubmitted(ctx context.Context, in *CancelSubmittedReq, opts ...grpc.CallOption) (*CancelSubmittedResp, error)
	CleanUpSubmitted(ctx context.Context, in *CleanUpSubmittedReq, opts ...grpc.CallOption) (*CleanUpSubmittedResp, error)
	DisposeSubmission(ctx context.Context, in *DisposeSubmissionReq, opts ...grpc.CallOption) (*DisposeSubmissionResp, error)
}

type commanderClient struct {
	cc grpc.ClientConnInterface
}

func NewCommanderClient(cc grpc.ClientConnInterface) CommanderClient {
	return &commanderClient{cc}
}

func (c *commanderClient) Handshake(ctx context.Context, in *HandshakeReq, opts ...grpc.CallOption) (*HandshakeResp, error) {
	out := new(HandshakeResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) ConnectFromAgent(ctx context.Context, in *ConnectFromAgentReq, opts ...grpc.CallOption) (*ConnectFromAgentResp, error) {
	out := new(ConnectFromAgentResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/ConnectFromAgent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) PreCheckScript(ctx context.Context, in *PreCheckScriptReq, opts ...grpc.CallOption) (*PreCheckScriptResp, error) {
	out := new(PreCheckScriptResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/PreCheckScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) PrepareScript(ctx context.Context, in *PrepareScriptReq, opts ...grpc.CallOption) (*PrepareScriptResp, error) {
	out := new(PrepareScriptResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/PrepareScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) SubmitScript(ctx context.Context, in *SubmitScriptReq, opts ...grpc.CallOption) (*SubmitScriptResp, error) {
	out := new(SubmitScriptResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/SubmitScript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) CancelSubmitted(ctx context.Context, in *CancelSubmittedReq, opts ...grpc.CallOption) (*CancelSubmittedResp, error) {
	out := new(CancelSubmittedResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/CancelSubmitted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) CleanUpSubmitted(ctx context.Context, in *CleanUpSubmittedReq, opts ...grpc.CallOption) (*CleanUpSubmittedResp, error) {
	out := new(CleanUpSubmittedResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/CleanUpSubmitted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commanderClient) DisposeSubmission(ctx context.Context, in *DisposeSubmissionReq, opts ...grpc.CallOption) (*DisposeSubmissionResp, error) {
	out := new(DisposeSubmissionResp)
	err := c.cc.Invoke(ctx, "/commander.Commander/DisposeSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommanderServer is the server API for Commander service.
// All implementations must embed UnimplementedCommanderServer
// for forward compatibility
type CommanderServer interface {
	// For handshake
	Handshake(context.Context, *HandshakeReq) (*HandshakeResp, error)
	Ping(context.Context, *PingReq) (*PingResp, error)
	ConnectFromAgent(context.Context, *ConnectFromAgentReq) (*ConnectFromAgentResp, error)
	// For execute task
	PreCheckScript(context.Context, *PreCheckScriptReq) (*PreCheckScriptResp, error)
	PrepareScript(context.Context, *PrepareScriptReq) (*PrepareScriptResp, error)
	SubmitScript(context.Context, *SubmitScriptReq) (*SubmitScriptResp, error)
	CancelSubmitted(context.Context, *CancelSubmittedReq) (*CancelSubmittedResp, error)
	CleanUpSubmitted(context.Context, *CleanUpSubmittedReq) (*CleanUpSubmittedResp, error)
	DisposeSubmission(context.Context, *DisposeSubmissionReq) (*DisposeSubmissionResp, error)
	mustEmbedUnimplementedCommanderServer()
}

// UnimplementedCommanderServer must be embedded to have forward compatible implementations.
type UnimplementedCommanderServer struct {
}

func (UnimplementedCommanderServer) Handshake(context.Context, *HandshakeReq) (*HandshakeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (UnimplementedCommanderServer) Ping(context.Context, *PingReq) (*PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCommanderServer) ConnectFromAgent(context.Context, *ConnectFromAgentReq) (*ConnectFromAgentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectFromAgent not implemented")
}
func (UnimplementedCommanderServer) PreCheckScript(context.Context, *PreCheckScriptReq) (*PreCheckScriptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreCheckScript not implemented")
}
func (UnimplementedCommanderServer) PrepareScript(context.Context, *PrepareScriptReq) (*PrepareScriptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareScript not implemented")
}
func (UnimplementedCommanderServer) SubmitScript(context.Context, *SubmitScriptReq) (*SubmitScriptResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScript not implemented")
}
func (UnimplementedCommanderServer) CancelSubmitted(context.Context, *CancelSubmittedReq) (*CancelSubmittedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSubmitted not implemented")
}
func (UnimplementedCommanderServer) CleanUpSubmitted(context.Context, *CleanUpSubmittedReq) (*CleanUpSubmittedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanUpSubmitted not implemented")
}
func (UnimplementedCommanderServer) DisposeSubmission(context.Context, *DisposeSubmissionReq) (*DisposeSubmissionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisposeSubmission not implemented")
}
func (UnimplementedCommanderServer) mustEmbedUnimplementedCommanderServer() {}

// UnsafeCommanderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommanderServer will
// result in compilation errors.
type UnsafeCommanderServer interface {
	mustEmbedUnimplementedCommanderServer()
}

func RegisterCommanderServer(s grpc.ServiceRegistrar, srv CommanderServer) {
	s.RegisterService(&Commander_ServiceDesc, srv)
}

func _Commander_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).Handshake(ctx, req.(*HandshakeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_ConnectFromAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectFromAgentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).ConnectFromAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/ConnectFromAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).ConnectFromAgent(ctx, req.(*ConnectFromAgentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_PreCheckScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreCheckScriptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).PreCheckScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/PreCheckScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).PreCheckScript(ctx, req.(*PreCheckScriptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_PrepareScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareScriptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).PrepareScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/PrepareScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).PrepareScript(ctx, req.(*PrepareScriptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_SubmitScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitScriptReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).SubmitScript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/SubmitScript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).SubmitScript(ctx, req.(*SubmitScriptReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_CancelSubmitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelSubmittedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).CancelSubmitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/CancelSubmitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).CancelSubmitted(ctx, req.(*CancelSubmittedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_CleanUpSubmitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanUpSubmittedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).CleanUpSubmitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/CleanUpSubmitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).CleanUpSubmitted(ctx, req.(*CleanUpSubmittedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commander_DisposeSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisposeSubmissionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommanderServer).DisposeSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.Commander/DisposeSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommanderServer).DisposeSubmission(ctx, req.(*DisposeSubmissionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Commander_ServiceDesc is the grpc.ServiceDesc for Commander service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Commander_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commander.Commander",
	HandlerType: (*CommanderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Commander_Handshake_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Commander_Ping_Handler,
		},
		{
			MethodName: "ConnectFromAgent",
			Handler:    _Commander_ConnectFromAgent_Handler,
		},
		{
			MethodName: "PreCheckScript",
			Handler:    _Commander_PreCheckScript_Handler,
		},
		{
			MethodName: "PrepareScript",
			Handler:    _Commander_PrepareScript_Handler,
		},
		{
			MethodName: "SubmitScript",
			Handler:    _Commander_SubmitScript_Handler,
		},
		{
			MethodName: "CancelSubmitted",
			Handler:    _Commander_CancelSubmitted_Handler,
		},
		{
			MethodName: "CleanUpSubmitted",
			Handler:    _Commander_CleanUpSubmitted_Handler,
		},
		{
			MethodName: "DisposeSubmission",
			Handler:    _Commander_DisposeSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commander.proto",
}

// AssistAgentClient is the client API for AssistAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssistAgentClient interface {
	// For handshake
	RegisterCommander(ctx context.Context, in *RegisterCommanderReq, opts ...grpc.CallOption) (*RegisterCommanderResp, error)
	// For execute task
	WriteSubmissionOutput(ctx context.Context, in *WriteSubmissionOutputReq, opts ...grpc.CallOption) (*WriteSubmissionOutputResp, error)
	FinalizeSubmissionStatus(ctx context.Context, in *FinalizeSubmissionStatusReq, opts ...grpc.CallOption) (*FinalizeSubmissionStatusResp, error)
}

type assistAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewAssistAgentClient(cc grpc.ClientConnInterface) AssistAgentClient {
	return &assistAgentClient{cc}
}

func (c *assistAgentClient) RegisterCommander(ctx context.Context, in *RegisterCommanderReq, opts ...grpc.CallOption) (*RegisterCommanderResp, error) {
	out := new(RegisterCommanderResp)
	err := c.cc.Invoke(ctx, "/commander.AssistAgent/RegisterCommander", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistAgentClient) WriteSubmissionOutput(ctx context.Context, in *WriteSubmissionOutputReq, opts ...grpc.CallOption) (*WriteSubmissionOutputResp, error) {
	out := new(WriteSubmissionOutputResp)
	err := c.cc.Invoke(ctx, "/commander.AssistAgent/WriteSubmissionOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assistAgentClient) FinalizeSubmissionStatus(ctx context.Context, in *FinalizeSubmissionStatusReq, opts ...grpc.CallOption) (*FinalizeSubmissionStatusResp, error) {
	out := new(FinalizeSubmissionStatusResp)
	err := c.cc.Invoke(ctx, "/commander.AssistAgent/FinalizeSubmissionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssistAgentServer is the server API for AssistAgent service.
// All implementations must embed UnimplementedAssistAgentServer
// for forward compatibility
type AssistAgentServer interface {
	// For handshake
	RegisterCommander(context.Context, *RegisterCommanderReq) (*RegisterCommanderResp, error)
	// For execute task
	WriteSubmissionOutput(context.Context, *WriteSubmissionOutputReq) (*WriteSubmissionOutputResp, error)
	FinalizeSubmissionStatus(context.Context, *FinalizeSubmissionStatusReq) (*FinalizeSubmissionStatusResp, error)
	mustEmbedUnimplementedAssistAgentServer()
}

// UnimplementedAssistAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAssistAgentServer struct {
}

func (UnimplementedAssistAgentServer) RegisterCommander(context.Context, *RegisterCommanderReq) (*RegisterCommanderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCommander not implemented")
}
func (UnimplementedAssistAgentServer) WriteSubmissionOutput(context.Context, *WriteSubmissionOutputReq) (*WriteSubmissionOutputResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteSubmissionOutput not implemented")
}
func (UnimplementedAssistAgentServer) FinalizeSubmissionStatus(context.Context, *FinalizeSubmissionStatusReq) (*FinalizeSubmissionStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeSubmissionStatus not implemented")
}
func (UnimplementedAssistAgentServer) mustEmbedUnimplementedAssistAgentServer() {}

// UnsafeAssistAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssistAgentServer will
// result in compilation errors.
type UnsafeAssistAgentServer interface {
	mustEmbedUnimplementedAssistAgentServer()
}

func RegisterAssistAgentServer(s grpc.ServiceRegistrar, srv AssistAgentServer) {
	s.RegisterService(&AssistAgent_ServiceDesc, srv)
}

func _AssistAgent_RegisterCommander_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCommanderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistAgentServer).RegisterCommander(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.AssistAgent/RegisterCommander",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistAgentServer).RegisterCommander(ctx, req.(*RegisterCommanderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistAgent_WriteSubmissionOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteSubmissionOutputReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistAgentServer).WriteSubmissionOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.AssistAgent/WriteSubmissionOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistAgentServer).WriteSubmissionOutput(ctx, req.(*WriteSubmissionOutputReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssistAgent_FinalizeSubmissionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeSubmissionStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssistAgentServer).FinalizeSubmissionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commander.AssistAgent/FinalizeSubmissionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssistAgentServer).FinalizeSubmissionStatus(ctx, req.(*FinalizeSubmissionStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AssistAgent_ServiceDesc is the grpc.ServiceDesc for AssistAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssistAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commander.AssistAgent",
	HandlerType: (*AssistAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCommander",
			Handler:    _AssistAgent_RegisterCommander_Handler,
		},
		{
			MethodName: "WriteSubmissionOutput",
			Handler:    _AssistAgent_WriteSubmissionOutput_Handler,
		},
		{
			MethodName: "FinalizeSubmissionStatus",
			Handler:    _AssistAgent_FinalizeSubmissionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commander.proto",
}
