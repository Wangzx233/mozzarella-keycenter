package token

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"mozzarella-keycenter/config"
	"mozzarella-keycenter/cryp"
	"mozzarella-keycenter/key"
	"strings"
	"time"
)

func CreateToken(Domain string, Uid string) (token string, rt string, exp int64, err error) {
	exp = time.Now().Add(config.DefaultTokenExpiresTime).Unix()
	//构建payload
	var payload = Payload{
		StandardPayload: StandardPayload{
			Audience:  config.DefaultTokenAudience,
			ExpiresAt: exp,
			Id:        "",
			IssuedAt:  time.Now().Unix(),
			Issuer:    config.DefaultTokenIssuer,
			NotBefore: time.Now().Add(config.DefaultTokenNotBefore).Unix(),
			Subject:   Domain,
		},
		Uid: Uid,
	}

	//序列化payload
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("token err : create token err :", err)

		return
	}
	//base64url编码
	pay := base64.URLEncoding.EncodeToString(bytesPayload)

	//计算出签名
	k, err := key.GetKey(Domain)
	if err != nil {
		log.Println("token err : create token err :", err)

		return
	}

	sig := cryp.EnCrypto(k.PrivateKey, bytesPayload)
	//base64url编码
	signatrue := base64.URLEncoding.EncodeToString(sig)
	//组合token
	token = strings.Join([]string{pay, signatrue}, ".")

	rt = CreateRefreshToken(bytesPayload, config.RefreshRefreshTokenExpiresTime)

	return
}
