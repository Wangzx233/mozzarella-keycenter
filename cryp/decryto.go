package cryp

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
)

// Sign 验签
func Sign(pubKey *rsa.PublicKey, payload, signature []byte) error {
	h := sha256.New()
	h.Write(payload)
	hash := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash, signature)
	if err != nil {
		err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash, signature)
	}
	return err
}
