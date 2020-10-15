package client

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/arisetransfer/arise/proto"
	"github.com/arisetransfer/arise/utils"
	"github.com/schollz/progressbar/v3"
	"google.golang.org/grpc"
)

func Reciever(code string) {
	addr, port := utils.GetIPAddrAndPort()
	conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
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
	file, err := client.Reciever(context.Background(), &proto.RecieverRequest{Code: code})
	if err != nil {
		log.Fatalf("Error:- %v", err)
		return
	}
	fmt.Println("FileName: ", file.Name, " Hash: ", file.Hash, " Size: ", file.Size)
	bar := progressbar.Default(file.Size)
	ack, err := client.SharePublicKey(context.Background(), &proto.PublicKey{Key: key.Bytes(), Code: code})
	if err != nil {
		log.Fatalf("Error:- %v", err)
		return
	}
	fmt.Println(ack.Message)
	senderInfo, err := client.GetSenderInfo(context.Background(), &proto.Code{Code: code})
	if err != nil {
		log.Printf("Error : %v", err)
	}
	fmt.Println("Receiving Data from ", senderInfo.Ip)
	aesEncryptionKey, err := client.GetEncryptionKey(context.Background(), &proto.Code{Code: code})
	if err != nil {
		log.Printf("Error : %v", err)
	}
	decryptedEncryptionKey, err := privateKey.Decrypt(nil, aesEncryptionKey.Key, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	recv := &proto.RecieverRequest{Code: code}
	stream, err := client.DataRecieve(context.Background(), recv)
	if err != nil {
		panic(err)
	}
	var fullfile []byte
	var finalKey [32]byte
	copy(finalKey[:], decryptedEncryptionKey)
	for {
		strm, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		decryptedContent, err := utils.Decrypt(strm.Content, &finalKey)
		bar.Add(len(decryptedContent))
		if err != nil {
			log.Printf("Error : %v", err)
			return
		}
		fullfile = append(fullfile, decryptedContent...)
	}
	err = ioutil.WriteFile(file.Name, fullfile, 0644)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	if utils.FileHash(file.Name) == file.Hash {
		fmt.Println("File Hash Verified!")
		return
	} else {
		fmt.Println("Hash Mismatch!")
	}
}
