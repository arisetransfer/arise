package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/iamstefin/arise-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6969", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	defer conn.Close()
	client := proto.NewAriseClient(conn)
	code, err := client.Sender(context.Background(), &proto.SenderRequest{Name: "hello.png", Hash: "1234567890"})
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	fmt.Println("Code: ", code.Code)
	stream, err := client.DataSend(context.Background())
	for i := 0; i < 10; i++ {
		time.Sleep(1*time.Second)
		if err := stream.Send(&proto.Chunk{Code:code.Code,Content:[]byte("Hello How are you ? "+strconv.Itoa(i))}); err != nil {
			log.Fatalf("%v",err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	fmt.Println(reply.Text)
}
