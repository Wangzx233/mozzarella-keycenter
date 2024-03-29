package token

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"mozzarella-keycenter/cryp"
	"mozzarella-keycenter/key"
	"strings"
	"time"
)

func VerifyToken(token string) (err error) {
	//先将%20替换为=
	token = strings.Replace(token, "%20", "=", -1)

	part := strings.Split(token, ".")
	if len(part) < 2 {
		err = errors.New("token err")
		return
	}
	if part[0] == "" || part[1] == "" {
		err = errors.New("token err")
		return
	}
	payload, err := base64.URLEncoding.DecodeString(part[0])
	if err != nil {
		log.Println("verifyToken err : ", err)
		return
	}

	//反序列化payload为了拿到domain
	var p Payload
	err = json.Unmarshal(payload, &p)
	if err != nil {
		log.Println("verifyToken err : ", err)
		return
	}

	if p.ExpiresAt < time.Now().Unix() {
		err = errors.New("token expired")
		return
	}
	//拿到签名
	signature, err := base64.URLEncoding.DecodeString(part[1])
	if err != nil {
		log.Println("verifyToken err : ", err)
		return
	}
	//判断domain中的key是否存在
	_, ok := key.AllKey[p.Subject]
	if !ok {
		log.Println("domain not exit", err)
		err = errors.New("domain not exit")
		return
	}
	//验证
	err = cryp.Sign(&key.AllKey[p.Subject].Now.Key.PublicKey, payload, signature)
	if err != nil {
		log.Println("verifyToken err : ", err)
		return
	}
	return
}
