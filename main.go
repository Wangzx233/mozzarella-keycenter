package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"mozzarella-keycenter/cryp"
	"os"
)

func main() {
	//1.打开私钥文件
	file, err := os.Open("private.pem")
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}
	privateKey := buf
	//1.打开公钥文件
	file, err = os.Open("public.pem")
	if err != nil {
		panic(err)
	}
	fileInfo, err = file.Stat()
	if err != nil {
		panic(err)
	}
	buf = make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		panic(err)
	}

	publicKey := buf
	fmt.Println()
	crypt := cryp.EnCrypto(privateKey, []byte("我是数据"))
	//var base []byte
	toString := base64.URLEncoding.EncodeToString(crypt)
	//fmt.Println(toString)

	bytes, err := base64.URLEncoding.DecodeString(toString)
	if err != nil {
		log.Println(err)
		return
	}

	h := sha256.New()
	h.Write([]byte("我是数据"))
	hash := h.Sum(nil)

	block, _ := pem.Decode(publicKey)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println("rpc check err:", err)
		return
	}

	pubkey := pubInterface.(*rsa.PublicKey)
	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hash, bytes)

	fmt.Println(err)
}
