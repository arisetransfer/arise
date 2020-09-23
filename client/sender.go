package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
	"encoding/hex"
	"crypto/sha256"

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
	file, err := os.Open("./"+os.Args[1])
	if err != nil {
		log.Println("File Not Found!")
		return
	}
	defer file.Close()
	f, err := os.Open("./"+os.Args[1])
	if err != nil {
		log.Println("File Not Found!")
		return
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	hash := hex.EncodeToString(h.Sum(nil))
	code, err := client.Sender(context.Background(), &proto.SenderRequest{Name: string(os.Args[1]), Hash: hash})
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
		if err := stream.Send(&proto.Chunk{Code:code.Code,Content:[]byte(buf[0:n])}); err != nil {
			log.Fatalf("%v",err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	fmt.Println(reply.Text)
}
