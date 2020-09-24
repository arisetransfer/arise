package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
	"github.com/iamstefin/arise-grpc/proto"
	"google.golang.org/grpc"
	"github.com/iamstefin/arise-grpc/utils"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6969", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	defer conn.Close()
	client := proto.NewAriseClient(conn)
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println("File Not Found!")
		return
	}
	defer file.Close()
	fname, _ := os.Stat(os.Args[1])
	code, err := client.Sender(context.Background(), &proto.SenderRequest{Name: fname.Name(), Hash: utils.FileHash(os.Args[1])})
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	fmt.Println("Code: ", code.Code)
	stream, err := client.DataSend(context.Background())
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	count := int(fname.Size()/1024) + 1
	bar := pb.StartNew(count)
	buf := make([]byte, 1024)
	reader := bufio.NewReader(file)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		if err := stream.Send(&proto.Chunk{Code: code.Code, Content: []byte(buf[0:n])}); err != nil {
			log.Fatalf("%v", err)
		}
		bar.Increment()
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	fmt.Println(reply.Text)
}
