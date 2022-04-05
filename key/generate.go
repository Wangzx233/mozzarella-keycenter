package key

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"mozzarella-keycenter/config"
	"sync"
	"time"
)

// Key RSA秘钥对
type Key struct {
	Key        *rsa.PrivateKey
	PrivateKey []byte
	PublicKey  []byte
}

// StorageKey RSA秘钥对
type StorageKey struct {
	//Old Key
	Now Key
	sync.RWMutex
}

// AllKey 内存中存放所有key
var AllKey = make(map[string]*StorageKey)

// AllKeyWLock AllKey写入锁
var AllKeyWLock sync.Mutex

func GetKey(domain string) (Key Key, err error) {
	storageKey, ok := AllKey[domain]
	if !ok {
		storageKey = &StorageKey{}
		go storageKey.GenerateKey()
		time.Sleep(time.Millisecond * 500)
		AllKey[domain] = storageKey
	}

	Key = storageKey.Now

	return
}

// GenerateKey 生成一对rsa秘钥并定时刷新
func (key *StorageKey) GenerateKey() {

	for true {
		key.Lock()
		k := generateKey()
		key.Now = Key{
			Key:        k,
			PublicKey:  GeneratePublicKey(&k.PublicKey),
			PrivateKey: GeneratePrivateKey(k),
		}
		key.Unlock()
		time.Sleep(config.RsaKeyRefreshTime)
	}
}

// generateKey 生成一对rsa秘钥
func generateKey() (k *rsa.PrivateKey) {
	pri, _ := rsa.GenerateKey(rand.Reader, 2048)
	return pri
}

// GeneratePrivateKey 从rsa秘钥对生成一个pem格式的私钥
func GeneratePrivateKey(pri *rsa.PrivateKey) (b []byte) {
	derStream := x509.MarshalPKCS1PrivateKey(pri)
	block := &pem.Block{
		Type:  "private key",
		Bytes: derStream,
	}
	var w bytes.Buffer
	_ = pem.Encode(&w, block)
	b = w.Bytes()
	return
}

// GeneratePublicKey 从rsa秘钥对生成一个pem格式的公钥
func GeneratePublicKey(pub *rsa.PublicKey) (b []byte) {
	derPkix, _ := x509.MarshalPKIXPublicKey(pub)
	block := &pem.Block{
		Type:  "public key",
		Bytes: derPkix,
	}
	var w bytes.Buffer
	_ = pem.Encode(&w, block)
	b = w.Bytes()

	return
}
