package cryp

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
)

// EnCrypto 加密 先取sha256散列 然后使用rsa私钥加密
func EnCrypto(key []byte, data []byte) (b []byte) {
	block, _ := pem.Decode(key)

	pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return
	}

	h := sha256.New()
	h.Write(data)
	hash := h.Sum(nil)

	b, err = rsa.SignPKCS1v15(rand.Reader, pri, crypto.SHA256, hash)
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
