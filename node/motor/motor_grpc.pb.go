// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package motor

import (
	context "context"
	api "github.com/sonr-io/core/node/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MotorStubClient is the client API for MotorStub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MotorStubClient interface {
	// Node Methods
	// Verification Method Request for Signed Data
	Edit(ctx context.Context, in *api.EditRequest, opts ...grpc.CallOption) (*api.EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(ctx context.Context, in *api.FetchRequest, opts ...grpc.CallOption) (*api.FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(ctx context.Context, in *api.ShareRequest, opts ...grpc.CallOption) (*api.ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(ctx context.Context, in *api.RespondRequest, opts ...grpc.CallOption) (*api.RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(ctx context.Context, in *api.SearchRequest, opts ...grpc.CallOption) (*api.SearchResponse, error)
	// Events Streams
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnLobbyRefreshClient, error)
	// Returns a stream of Mailbox Message Events
	OnMailboxMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnMailboxMessageClient, error)
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransmitAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitAcceptedClient, error)
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransmitDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitDeclinedClient, error)
	// Returns a stream of DecisionEvent's for Invites
	OnTransmitInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitInviteClient, error)
	// Returns a stream of ProgressEvent's for Sessions
	OnTransmitProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitProgressClient, error)
	// Returns a stream of Completed Transfers
	OnTransmitComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitCompleteClient, error)
}

type motorStubClient struct {
	cc grpc.ClientConnInterface
}

func NewMotorStubClient(cc grpc.ClientConnInterface) MotorStubClient {
	return &motorStubClient{cc}
}

func (c *motorStubClient) Edit(ctx context.Context, in *api.EditRequest, opts ...grpc.CallOption) (*api.EditResponse, error) {
	out := new(api.EditResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.MotorStub/Edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorStubClient) Fetch(ctx context.Context, in *api.FetchRequest, opts ...grpc.CallOption) (*api.FetchResponse, error) {
	out := new(api.FetchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.MotorStub/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorStubClient) Share(ctx context.Context, in *api.ShareRequest, opts ...grpc.CallOption) (*api.ShareResponse, error) {
	out := new(api.ShareResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.MotorStub/Share", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorStubClient) Respond(ctx context.Context, in *api.RespondRequest, opts ...grpc.CallOption) (*api.RespondResponse, error) {
	out := new(api.RespondResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.MotorStub/Respond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorStubClient) Search(ctx context.Context, in *api.SearchRequest, opts ...grpc.CallOption) (*api.SearchResponse, error) {
	out := new(api.SearchResponse)
	err := c.cc.Invoke(ctx, "/sonr.node.MotorStub/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *motorStubClient) OnLobbyRefresh(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnLobbyRefreshClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[0], "/sonr.node.MotorStub/OnLobbyRefresh", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnLobbyRefreshClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnLobbyRefreshClient interface {
	Recv() (*api.RefreshEvent, error)
	grpc.ClientStream
}

type motorStubOnLobbyRefreshClient struct {
	grpc.ClientStream
}

func (x *motorStubOnLobbyRefreshClient) Recv() (*api.RefreshEvent, error) {
	m := new(api.RefreshEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnMailboxMessage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnMailboxMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[1], "/sonr.node.MotorStub/OnMailboxMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnMailboxMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnMailboxMessageClient interface {
	Recv() (*api.MailboxEvent, error)
	grpc.ClientStream
}

type motorStubOnMailboxMessageClient struct {
	grpc.ClientStream
}

func (x *motorStubOnMailboxMessageClient) Recv() (*api.MailboxEvent, error) {
	m := new(api.MailboxEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnTransmitAccepted(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitAcceptedClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[2], "/sonr.node.MotorStub/OnTransmitAccepted", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnTransmitAcceptedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnTransmitAcceptedClient interface {
	Recv() (*api.DecisionEvent, error)
	grpc.ClientStream
}

type motorStubOnTransmitAcceptedClient struct {
	grpc.ClientStream
}

func (x *motorStubOnTransmitAcceptedClient) Recv() (*api.DecisionEvent, error) {
	m := new(api.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnTransmitDeclined(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitDeclinedClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[3], "/sonr.node.MotorStub/OnTransmitDeclined", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnTransmitDeclinedClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnTransmitDeclinedClient interface {
	Recv() (*api.DecisionEvent, error)
	grpc.ClientStream
}

type motorStubOnTransmitDeclinedClient struct {
	grpc.ClientStream
}

func (x *motorStubOnTransmitDeclinedClient) Recv() (*api.DecisionEvent, error) {
	m := new(api.DecisionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnTransmitInvite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitInviteClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[4], "/sonr.node.MotorStub/OnTransmitInvite", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnTransmitInviteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnTransmitInviteClient interface {
	Recv() (*api.InviteEvent, error)
	grpc.ClientStream
}

type motorStubOnTransmitInviteClient struct {
	grpc.ClientStream
}

func (x *motorStubOnTransmitInviteClient) Recv() (*api.InviteEvent, error) {
	m := new(api.InviteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnTransmitProgress(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitProgressClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[5], "/sonr.node.MotorStub/OnTransmitProgress", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnTransmitProgressClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnTransmitProgressClient interface {
	Recv() (*api.ProgressEvent, error)
	grpc.ClientStream
}

type motorStubOnTransmitProgressClient struct {
	grpc.ClientStream
}

func (x *motorStubOnTransmitProgressClient) Recv() (*api.ProgressEvent, error) {
	m := new(api.ProgressEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *motorStubClient) OnTransmitComplete(ctx context.Context, in *Empty, opts ...grpc.CallOption) (MotorStub_OnTransmitCompleteClient, error) {
	stream, err := c.cc.NewStream(ctx, &MotorStub_ServiceDesc.Streams[6], "/sonr.node.MotorStub/OnTransmitComplete", opts...)
	if err != nil {
		return nil, err
	}
	x := &motorStubOnTransmitCompleteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MotorStub_OnTransmitCompleteClient interface {
	Recv() (*api.CompleteEvent, error)
	grpc.ClientStream
}

type motorStubOnTransmitCompleteClient struct {
	grpc.ClientStream
}

func (x *motorStubOnTransmitCompleteClient) Recv() (*api.CompleteEvent, error) {
	m := new(api.CompleteEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MotorStubServer is the server API for MotorStub service.
// All implementations must embed UnimplementedMotorStubServer
// for forward compatibility
type MotorStubServer interface {
	// Node Methods
	// Verification Method Request for Signed Data
	Edit(context.Context, *api.EditRequest) (*api.EditResponse, error)
	// Fetch method finds data from Key/Value store
	Fetch(context.Context, *api.FetchRequest) (*api.FetchResponse, error)
	// Respond Method to an Invite with Decision
	Share(context.Context, *api.ShareRequest) (*api.ShareResponse, error)
	// Respond Method to an Invite with Decision
	Respond(context.Context, *api.RespondRequest) (*api.RespondResponse, error)
	// Search Method to find a Peer by SName or PeerID
	Search(context.Context, *api.SearchRequest) (*api.SearchResponse, error)
	// Events Streams
	// Returns a stream of Lobby Refresh Events
	OnLobbyRefresh(*Empty, MotorStub_OnLobbyRefreshServer) error
	// Returns a stream of Mailbox Message Events
	OnMailboxMessage(*Empty, MotorStub_OnMailboxMessageServer) error
	// Returns a stream of DecisionEvent's for Accepted Invites
	OnTransmitAccepted(*Empty, MotorStub_OnTransmitAcceptedServer) error
	// Returns a stream of DecisionEvent's for Rejected Invites
	OnTransmitDeclined(*Empty, MotorStub_OnTransmitDeclinedServer) error
	// Returns a stream of DecisionEvent's for Invites
	OnTransmitInvite(*Empty, MotorStub_OnTransmitInviteServer) error
	// Returns a stream of ProgressEvent's for Sessions
	OnTransmitProgress(*Empty, MotorStub_OnTransmitProgressServer) error
	// Returns a stream of Completed Transfers
	OnTransmitComplete(*Empty, MotorStub_OnTransmitCompleteServer) error
	mustEmbedUnimplementedMotorStubServer()
}

// UnimplementedMotorStubServer must be embedded to have forward compatible implementations.
type UnimplementedMotorStubServer struct {
}

func (UnimplementedMotorStubServer) Edit(context.Context, *api.EditRequest) (*api.EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedMotorStubServer) Fetch(context.Context, *api.FetchRequest) (*api.FetchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedMotorStubServer) Share(context.Context, *api.ShareRequest) (*api.ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Share not implemented")
}
func (UnimplementedMotorStubServer) Respond(context.Context, *api.RespondRequest) (*api.RespondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Respond not implemented")
}
func (UnimplementedMotorStubServer) Search(context.Context, *api.SearchRequest) (*api.SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedMotorStubServer) OnLobbyRefresh(*Empty, MotorStub_OnLobbyRefreshServer) error {
	return status.Errorf(codes.Unimplemented, "method OnLobbyRefresh not implemented")
}
func (UnimplementedMotorStubServer) OnMailboxMessage(*Empty, MotorStub_OnMailboxMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method OnMailboxMessage not implemented")
}
func (UnimplementedMotorStubServer) OnTransmitAccepted(*Empty, MotorStub_OnTransmitAcceptedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransmitAccepted not implemented")
}
func (UnimplementedMotorStubServer) OnTransmitDeclined(*Empty, MotorStub_OnTransmitDeclinedServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransmitDeclined not implemented")
}
func (UnimplementedMotorStubServer) OnTransmitInvite(*Empty, MotorStub_OnTransmitInviteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransmitInvite not implemented")
}
func (UnimplementedMotorStubServer) OnTransmitProgress(*Empty, MotorStub_OnTransmitProgressServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransmitProgress not implemented")
}
func (UnimplementedMotorStubServer) OnTransmitComplete(*Empty, MotorStub_OnTransmitCompleteServer) error {
	return status.Errorf(codes.Unimplemented, "method OnTransmitComplete not implemented")
}
func (UnimplementedMotorStubServer) mustEmbedUnimplementedMotorStubServer() {}

// UnsafeMotorStubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MotorStubServer will
// result in compilation errors.
type UnsafeMotorStubServer interface {
	mustEmbedUnimplementedMotorStubServer()
}

func RegisterMotorStubServer(s grpc.ServiceRegistrar, srv MotorStubServer) {
	s.RegisterService(&MotorStub_ServiceDesc, srv)
}

func _MotorStub_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorStubServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.MotorStub/Edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorStubServer).Edit(ctx, req.(*api.EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorStub_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorStubServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.MotorStub/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorStubServer).Fetch(ctx, req.(*api.FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorStub_Share_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.ShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorStubServer).Share(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.MotorStub/Share",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorStubServer).Share(ctx, req.(*api.ShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorStub_Respond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.RespondRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorStubServer).Respond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.MotorStub/Respond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorStubServer).Respond(ctx, req.(*api.RespondRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorStub_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MotorStubServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonr.node.MotorStub/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MotorStubServer).Search(ctx, req.(*api.SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MotorStub_OnLobbyRefresh_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnLobbyRefresh(m, &motorStubOnLobbyRefreshServer{stream})
}

type MotorStub_OnLobbyRefreshServer interface {
	Send(*api.RefreshEvent) error
	grpc.ServerStream
}

type motorStubOnLobbyRefreshServer struct {
	grpc.ServerStream
}

func (x *motorStubOnLobbyRefreshServer) Send(m *api.RefreshEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnMailboxMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnMailboxMessage(m, &motorStubOnMailboxMessageServer{stream})
}

type MotorStub_OnMailboxMessageServer interface {
	Send(*api.MailboxEvent) error
	grpc.ServerStream
}

type motorStubOnMailboxMessageServer struct {
	grpc.ServerStream
}

func (x *motorStubOnMailboxMessageServer) Send(m *api.MailboxEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnTransmitAccepted_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnTransmitAccepted(m, &motorStubOnTransmitAcceptedServer{stream})
}

type MotorStub_OnTransmitAcceptedServer interface {
	Send(*api.DecisionEvent) error
	grpc.ServerStream
}

type motorStubOnTransmitAcceptedServer struct {
	grpc.ServerStream
}

func (x *motorStubOnTransmitAcceptedServer) Send(m *api.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnTransmitDeclined_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnTransmitDeclined(m, &motorStubOnTransmitDeclinedServer{stream})
}

type MotorStub_OnTransmitDeclinedServer interface {
	Send(*api.DecisionEvent) error
	grpc.ServerStream
}

type motorStubOnTransmitDeclinedServer struct {
	grpc.ServerStream
}

func (x *motorStubOnTransmitDeclinedServer) Send(m *api.DecisionEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnTransmitInvite_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnTransmitInvite(m, &motorStubOnTransmitInviteServer{stream})
}

type MotorStub_OnTransmitInviteServer interface {
	Send(*api.InviteEvent) error
	grpc.ServerStream
}

type motorStubOnTransmitInviteServer struct {
	grpc.ServerStream
}

func (x *motorStubOnTransmitInviteServer) Send(m *api.InviteEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnTransmitProgress_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnTransmitProgress(m, &motorStubOnTransmitProgressServer{stream})
}

type MotorStub_OnTransmitProgressServer interface {
	Send(*api.ProgressEvent) error
	grpc.ServerStream
}

type motorStubOnTransmitProgressServer struct {
	grpc.ServerStream
}

func (x *motorStubOnTransmitProgressServer) Send(m *api.ProgressEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _MotorStub_OnTransmitComplete_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MotorStubServer).OnTransmitComplete(m, &motorStubOnTransmitCompleteServer{stream})
}

type MotorStub_OnTransmitCompleteServer interface {
	Send(*api.CompleteEvent) error
	grpc.ServerStream
}

type motorStubOnTransmitCompleteServer struct {
	grpc.ServerStream
}

func (x *motorStubOnTransmitCompleteServer) Send(m *api.CompleteEvent) error {
	return x.ServerStream.SendMsg(m)
}

// MotorStub_ServiceDesc is the grpc.ServiceDesc for MotorStub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MotorStub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonr.node.MotorStub",
	HandlerType: (*MotorStubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Edit",
			Handler:    _MotorStub_Edit_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _MotorStub_Fetch_Handler,
		},
		{
			MethodName: "Share",
			Handler:    _MotorStub_Share_Handler,
		},
		{
			MethodName: "Respond",
			Handler:    _MotorStub_Respond_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _MotorStub_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnLobbyRefresh",
			Handler:       _MotorStub_OnLobbyRefresh_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnMailboxMessage",
			Handler:       _MotorStub_OnMailboxMessage_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransmitAccepted",
			Handler:       _MotorStub_OnTransmitAccepted_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransmitDeclined",
			Handler:       _MotorStub_OnTransmitDeclined_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransmitInvite",
			Handler:       _MotorStub_OnTransmitInvite_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransmitProgress",
			Handler:       _MotorStub_OnTransmitProgress_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnTransmitComplete",
			Handler:       _MotorStub_OnTransmitComplete_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/node/motor.proto",
}