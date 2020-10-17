package server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/arisetransfer/arise/proto"
	"github.com/arisetransfer/arise/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Server struct{}

var (
	connections         = make(map[string]proto.SenderRequest)
	contents            = make(map[string](chan []byte))
	done                = make(map[string]chan bool)
	dataSent            = make(map[string]chan bool)
	rip                 = make(map[string]proto.RecieverInfo)
	sip                 = make(map[string]proto.SenderInfo)
	recieverPublicKey   = make(map[string][]byte)
	senderEncryptionKey = make(map[string][]byte)
)

func StartRelay() {
	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	grpcServer := grpc.NewServer()

	server := &Server{}

	proto.RegisterAriseService(grpcServer, &proto.AriseService{
		Sender:             server.Sender,
		Reciever:           server.Reciever,
		DataSend:           server.DataSend,
		DataRecieve:        server.DataRecieve,
		GetRecieverInfo:    server.GetRecieverInfo,
		GetSenderInfo:      server.GetSenderInfo,
		GetPublicKey:       server.GetPublicKey,
		SharePublicKey:     server.SharePublicKey,
		GetEncryptionKey:   server.GetEncryptionKey,
		ShareEncryptionKey: server.ShareEncryptionKey,
	})

	fmt.Println("gRPC Server Started")
	grpcServer.Serve(lis)
}

func (s *Server) Sender(ctx context.Context, request *proto.SenderRequest) (*proto.SenderResponse, error) {
	for {
		code := utils.Dice(4)
		if _, ok := connections[code]; ok {
		} else {
			connections[code] = *request
			contents[code] = make(chan []byte)
			done[code] = make(chan bool, 1)
			dataSent[code] = make(chan bool, 1)
			senderIp, _ := peer.FromContext(ctx)
			sip[code] = proto.SenderInfo{Ip: senderIp.Addr.String()}
			go func() {
				//Deleting all the existing data upon Data Send or timeout!
				select {
				case <-done[code]:
					break
				case <-time.After(10 * time.Minute):
					break
				}
				delete(contents, code)
				delete(rip, code)
				delete(sip, code)
				delete(recieverPublicKey, code)
				delete(senderEncryptionKey, code)
				delete(connections, code)
				delete(recieverPublicKey, code)
				delete(senderEncryptionKey, code)
				delete(dataSent, code)
			}()
			log.Println("Recieving from ", senderIp.Addr.String())
			return &proto.SenderResponse{Code: code}, nil
		}
	}

	return &proto.SenderResponse{Code: ""}, errors.New("Cannot Generate Code")
}

func (s *Server) Reciever(ctx context.Context, request *proto.RecieverRequest) (*proto.RecieverResponse, error) {

	if _, ok := connections[request.Code]; ok {
		recieverIp, _ := peer.FromContext(ctx)
		rip[request.Code] = proto.RecieverInfo{Ip: recieverIp.Addr.String()}
		log.Println("Sending to ", recieverIp.Addr.String())
		return &proto.RecieverResponse{Name: connections[request.Code].Name, Hash: connections[request.Code].Hash, Size: connections[request.Code].Size}, nil
	}
	return &proto.RecieverResponse{Name: "", Hash: ""}, errors.New("The Code Is Invalid")
}

func (s *Server) DataSend(stream proto.Arise_DataSendServer) error {
	var code string
	for {
		data, err := stream.Recv()
		if err == nil {
			code = data.Code
		}
		if err == io.EOF {
			dataSent[code] <- true
			done[code] <- true
			return stream.SendAndClose(&proto.SendResponse{Text: "Data Sent Successfully!"})
		}
		if err != nil {
			log.Println("Error : ", err)
			dataSent[code] <- true
			done[code] <- true
			return stream.SendAndClose(&proto.SendResponse{Text: "Data Not Recieved!"})
		}
		contents[data.Code] <- data.Content
	}
}

func (s *Server) DataRecieve(request *proto.RecieverRequest, stream proto.Arise_DataRecieveServer) error {
Recieve:
	for {
		select {
		case content := <-contents[request.Code]:
			if err := stream.Send(&proto.RecieveResponse{Content: content}); err != nil {
				return err
			}
		case <-dataSent[request.Code]:
			break Recieve
		}
	}
	return nil
}

func (s *Server) GetRecieverInfo(ctx context.Context, request *proto.Code) (*proto.RecieverInfo, error) {
	for {
		time.Sleep(1 * time.Second)
		if val, ok := rip[request.Code]; ok {
			return &proto.RecieverInfo{Ip: val.Ip}, nil
		}
	}
}

func (s *Server) GetSenderInfo(ctx context.Context, request *proto.Code) (*proto.SenderInfo, error) {
	for {
		time.Sleep(1 * time.Second)
		if val, ok := sip[request.Code]; ok {
			return &proto.SenderInfo{Ip: val.Ip}, nil
		}
	}
}

func (s *Server) GetPublicKey(ctx context.Context, request *proto.Code) (*proto.PublicKey, error) {
	for {
		time.Sleep(1 * time.Second)
		if val, ok := recieverPublicKey[request.Code]; ok {
			return &proto.PublicKey{Key: val}, nil
		}
	}
}

func (s *Server) SharePublicKey(ctx context.Context, request *proto.PublicKey) (*proto.PublicKeyResponse, error) {
	recieverPublicKey[request.Code] = request.Key
	return &proto.PublicKeyResponse{Message: "PublicKey Recieved!"}, nil
}

func (s *Server) ShareEncryptionKey(ctx context.Context, request *proto.EncryptionKey) (*proto.EncryptionKeyResponse, error) {
	senderEncryptionKey[request.Code] = request.Key
	return &proto.EncryptionKeyResponse{Message: "EncryptionKey Recieved!"}, nil
}

func (s *Server) GetEncryptionKey(ctx context.Context, request *proto.Code) (*proto.EncryptionKey, error) {
	for {
		time.Sleep(1 * time.Second)
		if val, ok := senderEncryptionKey[request.Code]; ok {
			return &proto.EncryptionKey{Key: val}, nil
		}
	}
}
