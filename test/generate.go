package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mozzarella-keycenter/pb"
	"os"
)

func main() {
	//GenerateRsaKey(2048)

	//t, err := token.CreateToken("123", "111")
	//if err != nil {
	//	log.Println(err)
	//
	//	return
	//}
	//err = token.VerifyToken(t)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	// 连接服务器
	conn, err := grpc.Dial("8.142.81.74:8901", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewMozzarellaBookClient(conn)
	token, err := c.CreateToken(context.Background(), &pb.CreateTokenReq{
		Domain: "123",
		Uid:    "11",
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	// 调用服务端的SayHello
	r, err := c.VerifyToken(context.Background(), &pb.VerifyTokenReq{Token: token.Token + "123"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}

// GenerateRsaKey 生成rsa私钥和公钥并写入磁盘文件
func GenerateRsaKey(keySize int) (err error) {
	//1.生成rsa秘钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	//2.通过x509标准将得到的rsa私钥序列化为ASN.1的DER编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	//3.创建一个pem.Block结构体
	block := pem.Block{
		Type:  "rsa private key",
		Bytes: derText,
	}
	//4.通过pem将设置好的私钥数据进行编码，并写入磁盘文件
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	var w bytes.Buffer
	err = pem.Encode(&w, &block)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(w.Bytes())
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			log.Println(err)
			return
		}
	}
	w.Reset()

	// ==========公钥==================
	//1.从私钥中取出公钥
	publicKey := privateKey.PublicKey
	//2.使用x509序列化公钥为字符串
	marshalPKIXPublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//3.通过公钥字符串设置到pem格式块中
	block = pem.Block{
		Type:    "rsa public key",
		Headers: nil,
		Bytes:   marshalPKIXPublicKey,
	}
	//4.pem编码
	file, err = os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(&w, &block)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(w.Bytes())
	if err != nil {
		if err == io.EOF {
			err = nil
		} else {
			log.Println(err)
			return
		}
	}
	file.Close()
	return
}
