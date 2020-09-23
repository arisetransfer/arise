package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"io"
	"io/ioutil"
	"crypto/sha256"
	"encoding/hex"

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
	var fullfile []byte
	for {
		strm,err := stream.Recv()
		if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
		fullfile =  append(fullfile,strm.Content...)
	}
	err = ioutil.WriteFile(code.Name, fullfile, 0644)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	f, err := os.Open("./"+code.Name)
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
	if hash == code.Hash {
		fmt.Println("File Hash Verified!")
		return
	}else{
		fmt.Println("Hash Mismatch!")
	}
}
