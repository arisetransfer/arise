package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
  "io"

	"github.com/iamstefin/arise-grpc/proto"
	"github.com/iamstefin/arise-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Server struct{}

var (
	connections = make(map[string]proto.SenderRequest)
  contents = make(map[string](chan []byte))
  done = make(map[string]chan bool)
	rip = make(map[string]proto.RecieverInfo)
	sip = make(map[string]proto.SenderInfo)
)

func main() {
	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	grpcServer := grpc.NewServer()

	server := &Server{}

	proto.RegisterAriseService(grpcServer, &proto.AriseService{
		Sender:   server.Sender,
		Reciever: server.Reciever,
		DataSend: server.DataSend,
		DataRecieve: server.DataRecieve,
		GetRecieverInfo: server.GetRecieverInfo,
		GetSenderInfo: server.GetSenderInfo,
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
			senderIp,_ := peer.FromContext(ctx)
			sip[code] = proto.SenderInfo{Ip:senderIp.Addr.String()}
			log.Println("Recieving from ",senderIp.Addr.String())
			return &proto.SenderResponse{Code: code}, nil
		}
	}

	return &proto.SenderResponse{Code: ""}, errors.New("Cannot Generate Code")
}

func (s *Server) Reciever(ctx context.Context, request *proto.RecieverRequest) (*proto.RecieverResponse, error) {

  if _, ok := connections[request.Code]; ok {
    defer delete(connections, request.Code)
		recieverIp,_ := peer.FromContext(ctx)
		rip[request.Code] = proto.RecieverInfo{Ip:recieverIp.Addr.String()}
		log.Println("Sending to ",recieverIp.Addr.String())
		return &proto.RecieverResponse{Name: connections[request.Code].Name, Hash: connections[request.Code].Hash}, nil
	}
	return &proto.RecieverResponse{Name: "", Hash: ""}, errors.New("The Code Is Invalid")
}


func (s *Server) DataSend(stream proto.Arise_DataSendServer) error {
  var code string
  for {
    data,err := stream.Recv()
    if err == nil{
      code = data.Code
    }
    if err == io.EOF {
      done[code]<-true
      return stream.SendAndClose(&proto.SendResponse{Text:"Data Sent Successfully!"})
    }
    if err != nil {
      log.Println("Error : ",err)
			done[code]<-true
			return stream.SendAndClose(&proto.SendResponse{Text:"Data Not Recieved!"})
    }
    contents[data.Code]<-data.Content
  }
}


func (s *Server) DataRecieve(request *proto.RecieverRequest,stream proto.Arise_DataRecieveServer) error {
  defer delete(contents,request.Code)
  Recieve:
  for {
    select {
    case content := <-contents[request.Code]:
      if err := stream.Send(&proto.RecieveResponse{Content:content}); err != nil {
        return err
      }
    case <- done[request.Code]:
      break Recieve
    }
  }
  return nil
}

func (s *Server) GetRecieverInfo(ctx context.Context,request *proto.Code) (*proto.RecieverInfo,error) {
	for {
		if(rip[request.Code].Ip!=""){
			return &proto.RecieverInfo{Ip:rip[request.Code].Ip},nil
		}
	}
}

func (s *Server) GetSenderInfo(ctx context.Context,request *proto.Code) (*proto.SenderInfo,error) {
	for {
		if(sip[request.Code].Ip!=""){
			return &proto.SenderInfo{Ip:sip[request.Code].Ip},nil
		}
	}
}
