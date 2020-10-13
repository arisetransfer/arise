package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"crypto/rsa"
	"crypto/rand"
	"bytes"
	"encoding/gob"
	"crypto"

	"github.com/arisetransfer/arise/proto"
	"google.golang.org/grpc"
	"github.com/arisetransfer/arise/utils"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:6969", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error:- ", err)
		return
	}
	defer conn.Close()
	client := proto.NewAriseClient(conn)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	var key bytes.Buffer
  enc := gob.NewEncoder(&key)
  err = enc.Encode(privateKey.PublicKey)
  if err != nil {
    panic(err)
  }
	ack,err := client.SharePublicKey(context.Background(),&proto.PublicKey{Key:key.Bytes(),Code:os.Args[1]})
	if err != nil {
		log.Fatalf("Error:- %v", err)
		return
	}
	fmt.Println(ack.Message)
	code, err := client.Reciever(context.Background(), &proto.RecieverRequest{Code: os.Args[1]})
	if err != nil {
		log.Fatalf("Error:- %v", err)
		return
	}
	fmt.Println("FileName: ", code.Name, " Hash: ", code.Hash)
	senderInfo,err := client.GetSenderInfo(context.Background(),&proto.Code{Code:os.Args[1]})
	if err != nil {
		log.Printf("Error : %v", err)
	}
	fmt.Println("Receiving Data from ",senderInfo.Ip)
	aesEncryptionKey,err := client.GetEncryptionKey(context.Background(),&proto.Code{Code:os.Args[1]})
	if err != nil {
		log.Printf("Error : %v", err)
	}
	decryptedEncryptionKey, err := privateKey.Decrypt(nil, aesEncryptionKey.Key, &rsa.OAEPOptions{Hash: crypto.SHA256})
  if err != nil {
	   panic(err)
   }
	recv := &proto.RecieverRequest{Code: os.Args[1]}
	stream, err := client.DataRecieve(context.Background(), recv)
	if err != nil {
		panic(err)
	}
	var fullfile []byte
	for {
		strm, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		var key [32]byte
		copy(key[:], decryptedEncryptionKey)
		decryptedContent,err := utils.Decrypt(strm.Content,&key)
		if err != nil {
			log.Printf("Error : %v", err)
			return
		}
		fullfile = append(fullfile, decryptedContent...)
	}
	err = ioutil.WriteFile(code.Name, fullfile, 0644)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	if utils.FileHash(code.Name) == code.Hash {
		fmt.Println("File Hash Verified!")
		return
	} else {
		fmt.Println("Hash Mismatch!")
	}
}
