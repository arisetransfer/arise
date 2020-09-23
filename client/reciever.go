package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"io"

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
	code, err := client.Reciever(context.Background(), &proto.RecieverRequest{Code: os.Args[1]})
	if err != nil {
		log.Fatalf("Error:- %v", err)
		return
	}
	fmt.Println("FileName: ", code.Name, " Hash: ", code.Hash)
	recv := &proto.RecieverRequest{Code:os.Args[1]}
	stream, err := client.DataRecieve(context.Background(), recv)
	if err != nil {
		panic(err)
	}
	for {
		strm,err := stream.Recv()
		if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
		log.Println(string(strm.Content))
	}
}
