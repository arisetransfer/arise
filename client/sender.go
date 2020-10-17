package client

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/arisetransfer/arise/proto"
	"github.com/arisetransfer/arise/utils"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/grpc"
)

func Sender(filename string) {
	addr, port := utils.GetIPAddrAndPort()
	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	defer conn.Close()
	client := proto.NewAriseClient(conn)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File Not Found!")
		return
	}
	defer file.Close()
	fname, _ := os.Stat(filename)
	bar := progressbar.Default(fname.Size())
	code, err := client.Sender(context.Background(), &proto.SenderRequest{Name: fname.Name(), Hash: utils.FileHash(filename), Size: fname.Size()})
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	fmt.Printf("On the other device type \n\n")
	fmt.Println("arise recieve", code.Code)
	recieverInfo, err := client.GetRecieverInfo(context.Background(), &proto.Code{Code: code.Code})
	if err != nil {
		log.Printf("Errror : %v", err)
	}
	fmt.Println("Sending Data to ", recieverInfo.Ip)
	stream, err := client.DataSend(context.Background())
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	publicKey, err := client.GetPublicKey(context.Background(), &proto.Code{Code: code.Code})
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	//fmt.Println(publicKey.Key)
	dec := gob.NewDecoder(bytes.NewBuffer(publicKey.Key))
	var decodedPublicKey rsa.PublicKey
	err = dec.Decode(&decodedPublicKey)
	if err != nil {
		log.Println(err)
	}
	aesEncryptionKey := utils.NewEncryptionKey()
	encryptedKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&decodedPublicKey,
		aesEncryptionKey[:32],
		nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.ShareEncryptionKey(context.Background(), &proto.EncryptionKey{Key: encryptedKey, Code: code.Code})
	if err != nil {
		log.Printf("Error :%v", err)
		return
	}
	fmt.Println(resp.Message)
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
		bar.Add(n)
		encryptedContent, err := utils.Encrypt([]byte(buf[0:n]), aesEncryptionKey)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		if err := stream.Send(&proto.Chunk{Code: code.Code, Content: encryptedContent}); err != nil {
			log.Fatalf("%v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	fmt.Println(reply.Text)
}
