// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AriseClient is the client API for Arise service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AriseClient interface {
	Sender(ctx context.Context, in *SenderRequest, opts ...grpc.CallOption) (*SenderResponse, error)
	Reciever(ctx context.Context, in *RecieverRequest, opts ...grpc.CallOption) (*RecieverResponse, error)
	DataSend(ctx context.Context, opts ...grpc.CallOption) (Arise_DataSendClient, error)
	DataRecieve(ctx context.Context, in *RecieverRequest, opts ...grpc.CallOption) (Arise_DataRecieveClient, error)
	GetRecieverInfo(ctx context.Context, in *Code, opts ...grpc.CallOption) (*RecieverInfo, error)
	GetSenderInfo(ctx context.Context, in *Code, opts ...grpc.CallOption) (*SenderInfo, error)
	GetPublicKey(ctx context.Context, in *Code, opts ...grpc.CallOption) (*PublicKey, error)
	SharePublicKey(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*PublicKeyResponse, error)
	GetEncryptionKey(ctx context.Context, in *Code, opts ...grpc.CallOption) (*EncryptionKey, error)
	ShareEncryptionKey(ctx context.Context, in *EncryptionKey, opts ...grpc.CallOption) (*EncryptionKeyResponse, error)
}

type ariseClient struct {
	cc grpc.ClientConnInterface
}

func NewAriseClient(cc grpc.ClientConnInterface) AriseClient {
	return &ariseClient{cc}
}

var ariseSenderStreamDesc = &grpc.StreamDesc{
	StreamName: "Sender",
}

func (c *ariseClient) Sender(ctx context.Context, in *SenderRequest, opts ...grpc.CallOption) (*SenderResponse, error) {
	out := new(SenderResponse)
	err := c.cc.Invoke(ctx, "/arise.Arise/Sender", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseRecieverStreamDesc = &grpc.StreamDesc{
	StreamName: "Reciever",
}

func (c *ariseClient) Reciever(ctx context.Context, in *RecieverRequest, opts ...grpc.CallOption) (*RecieverResponse, error) {
	out := new(RecieverResponse)
	err := c.cc.Invoke(ctx, "/arise.Arise/Reciever", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseDataSendStreamDesc = &grpc.StreamDesc{
	StreamName:    "DataSend",
	ClientStreams: true,
}

func (c *ariseClient) DataSend(ctx context.Context, opts ...grpc.CallOption) (Arise_DataSendClient, error) {
	stream, err := c.cc.NewStream(ctx, ariseDataSendStreamDesc, "/arise.Arise/DataSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &ariseDataSendClient{stream}
	return x, nil
}

type Arise_DataSendClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*SendResponse, error)
	grpc.ClientStream
}

type ariseDataSendClient struct {
	grpc.ClientStream
}

func (x *ariseDataSendClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ariseDataSendClient) CloseAndRecv() (*SendResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var ariseDataRecieveStreamDesc = &grpc.StreamDesc{
	StreamName:    "DataRecieve",
	ServerStreams: true,
}

func (c *ariseClient) DataRecieve(ctx context.Context, in *RecieverRequest, opts ...grpc.CallOption) (Arise_DataRecieveClient, error) {
	stream, err := c.cc.NewStream(ctx, ariseDataRecieveStreamDesc, "/arise.Arise/DataRecieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &ariseDataRecieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Arise_DataRecieveClient interface {
	Recv() (*RecieveResponse, error)
	grpc.ClientStream
}

type ariseDataRecieveClient struct {
	grpc.ClientStream
}

func (x *ariseDataRecieveClient) Recv() (*RecieveResponse, error) {
	m := new(RecieveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var ariseGetRecieverInfoStreamDesc = &grpc.StreamDesc{
	StreamName: "GetRecieverInfo",
}

func (c *ariseClient) GetRecieverInfo(ctx context.Context, in *Code, opts ...grpc.CallOption) (*RecieverInfo, error) {
	out := new(RecieverInfo)
	err := c.cc.Invoke(ctx, "/arise.Arise/GetRecieverInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseGetSenderInfoStreamDesc = &grpc.StreamDesc{
	StreamName: "GetSenderInfo",
}

func (c *ariseClient) GetSenderInfo(ctx context.Context, in *Code, opts ...grpc.CallOption) (*SenderInfo, error) {
	out := new(SenderInfo)
	err := c.cc.Invoke(ctx, "/arise.Arise/GetSenderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseGetPublicKeyStreamDesc = &grpc.StreamDesc{
	StreamName: "GetPublicKey",
}

func (c *ariseClient) GetPublicKey(ctx context.Context, in *Code, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := c.cc.Invoke(ctx, "/arise.Arise/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseSharePublicKeyStreamDesc = &grpc.StreamDesc{
	StreamName: "SharePublicKey",
}

func (c *ariseClient) SharePublicKey(ctx context.Context, in *PublicKey, opts ...grpc.CallOption) (*PublicKeyResponse, error) {
	out := new(PublicKeyResponse)
	err := c.cc.Invoke(ctx, "/arise.Arise/SharePublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseGetEncryptionKeyStreamDesc = &grpc.StreamDesc{
	StreamName: "GetEncryptionKey",
}

func (c *ariseClient) GetEncryptionKey(ctx context.Context, in *Code, opts ...grpc.CallOption) (*EncryptionKey, error) {
	out := new(EncryptionKey)
	err := c.cc.Invoke(ctx, "/arise.Arise/GetEncryptionKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var ariseShareEncryptionKeyStreamDesc = &grpc.StreamDesc{
	StreamName: "ShareEncryptionKey",
}

func (c *ariseClient) ShareEncryptionKey(ctx context.Context, in *EncryptionKey, opts ...grpc.CallOption) (*EncryptionKeyResponse, error) {
	out := new(EncryptionKeyResponse)
	err := c.cc.Invoke(ctx, "/arise.Arise/ShareEncryptionKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AriseService is the service API for Arise service.
// Fields should be assigned to their respective handler implementations only before
// RegisterAriseService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type AriseService struct {
	Sender             func(context.Context, *SenderRequest) (*SenderResponse, error)
	Reciever           func(context.Context, *RecieverRequest) (*RecieverResponse, error)
	DataSend           func(Arise_DataSendServer) error
	DataRecieve        func(*RecieverRequest, Arise_DataRecieveServer) error
	GetRecieverInfo    func(context.Context, *Code) (*RecieverInfo, error)
	GetSenderInfo      func(context.Context, *Code) (*SenderInfo, error)
	GetPublicKey       func(context.Context, *Code) (*PublicKey, error)
	SharePublicKey     func(context.Context, *PublicKey) (*PublicKeyResponse, error)
	GetEncryptionKey   func(context.Context, *Code) (*EncryptionKey, error)
	ShareEncryptionKey func(context.Context, *EncryptionKey) (*EncryptionKeyResponse, error)
}

func (s *AriseService) sender(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Sender(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/Sender",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Sender(ctx, req.(*SenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) reciever(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecieverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Reciever(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/Reciever",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Reciever(ctx, req.(*RecieverRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) dataSend(_ interface{}, stream grpc.ServerStream) error {
	return s.DataSend(&ariseDataSendServer{stream})
}
func (s *AriseService) dataRecieve(_ interface{}, stream grpc.ServerStream) error {
	m := new(RecieverRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.DataRecieve(m, &ariseDataRecieveServer{stream})
}
func (s *AriseService) getRecieverInfo(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetRecieverInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/GetRecieverInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetRecieverInfo(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) getSenderInfo(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetSenderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/GetSenderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSenderInfo(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) getPublicKey(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetPublicKey(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) sharePublicKey(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SharePublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/SharePublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SharePublicKey(ctx, req.(*PublicKey))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) getEncryptionKey(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Code)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetEncryptionKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/GetEncryptionKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetEncryptionKey(ctx, req.(*Code))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *AriseService) shareEncryptionKey(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptionKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ShareEncryptionKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/arise.Arise/ShareEncryptionKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ShareEncryptionKey(ctx, req.(*EncryptionKey))
	}
	return interceptor(ctx, in, info, handler)
}

type Arise_DataSendServer interface {
	SendAndClose(*SendResponse) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type ariseDataSendServer struct {
	grpc.ServerStream
}

func (x *ariseDataSendServer) SendAndClose(m *SendResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ariseDataSendServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type Arise_DataRecieveServer interface {
	Send(*RecieveResponse) error
	grpc.ServerStream
}

type ariseDataRecieveServer struct {
	grpc.ServerStream
}

func (x *ariseDataRecieveServer) Send(m *RecieveResponse) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterAriseService registers a service implementation with a gRPC server.
func RegisterAriseService(s grpc.ServiceRegistrar, srv *AriseService) {
	srvCopy := *srv
	if srvCopy.Sender == nil {
		srvCopy.Sender = func(context.Context, *SenderRequest) (*SenderResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Sender not implemented")
		}
	}
	if srvCopy.Reciever == nil {
		srvCopy.Reciever = func(context.Context, *RecieverRequest) (*RecieverResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Reciever not implemented")
		}
	}
	if srvCopy.DataSend == nil {
		srvCopy.DataSend = func(Arise_DataSendServer) error {
			return status.Errorf(codes.Unimplemented, "method DataSend not implemented")
		}
	}
	if srvCopy.DataRecieve == nil {
		srvCopy.DataRecieve = func(*RecieverRequest, Arise_DataRecieveServer) error {
			return status.Errorf(codes.Unimplemented, "method DataRecieve not implemented")
		}
	}
	if srvCopy.GetRecieverInfo == nil {
		srvCopy.GetRecieverInfo = func(context.Context, *Code) (*RecieverInfo, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetRecieverInfo not implemented")
		}
	}
	if srvCopy.GetSenderInfo == nil {
		srvCopy.GetSenderInfo = func(context.Context, *Code) (*SenderInfo, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetSenderInfo not implemented")
		}
	}
	if srvCopy.GetPublicKey == nil {
		srvCopy.GetPublicKey = func(context.Context, *Code) (*PublicKey, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
		}
	}
	if srvCopy.SharePublicKey == nil {
		srvCopy.SharePublicKey = func(context.Context, *PublicKey) (*PublicKeyResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method SharePublicKey not implemented")
		}
	}
	if srvCopy.GetEncryptionKey == nil {
		srvCopy.GetEncryptionKey = func(context.Context, *Code) (*EncryptionKey, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetEncryptionKey not implemented")
		}
	}
	if srvCopy.ShareEncryptionKey == nil {
		srvCopy.ShareEncryptionKey = func(context.Context, *EncryptionKey) (*EncryptionKeyResponse, error) {
			return nil, status.Errorf(codes.Unimplemented, "method ShareEncryptionKey not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "arise.Arise",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Sender",
				Handler:    srvCopy.sender,
			},
			{
				MethodName: "Reciever",
				Handler:    srvCopy.reciever,
			},
			{
				MethodName: "GetRecieverInfo",
				Handler:    srvCopy.getRecieverInfo,
			},
			{
				MethodName: "GetSenderInfo",
				Handler:    srvCopy.getSenderInfo,
			},
			{
				MethodName: "GetPublicKey",
				Handler:    srvCopy.getPublicKey,
			},
			{
				MethodName: "SharePublicKey",
				Handler:    srvCopy.sharePublicKey,
			},
			{
				MethodName: "GetEncryptionKey",
				Handler:    srvCopy.getEncryptionKey,
			},
			{
				MethodName: "ShareEncryptionKey",
				Handler:    srvCopy.shareEncryptionKey,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "DataSend",
				Handler:       srvCopy.dataSend,
				ClientStreams: true,
			},
			{
				StreamName:    "DataRecieve",
				Handler:       srvCopy.dataRecieve,
				ServerStreams: true,
			},
		},
		Metadata: "arise.proto",
	}

	s.RegisterService(&sd, nil)
}
