package token

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"mozzarella-keycenter/redisdao"
	"time"
)

// salt 盐
func salt() (pre []byte, fix []byte) {
	pre = make([]byte, 10)
	fix = make([]byte, 10)
	arr := [][]byte{pre, fix}
	for z := 0; z < 2; z++ {
		for i := 0; i < 10; i++ {
			arr[z][i] = byte(rand.Intn(255))
		}
	}
	return
}

// CreateRefreshToken 生成一个refresh_token
// 生成规则: 把payload拿出来 前后加盐 然后sha256 取16进制
func CreateRefreshToken(payload []byte, duration time.Duration) (rt string) {

	// 生成refresh_token
	pre, fix := salt()
	byte := bytes.Join([][]byte{pre, payload, fix}, nil)
	sha := sha256.New()
	sha.Write(byte)
	res := sha.Sum(nil)
	// refresh_token
	rt = fmt.Sprintf("%x", res)

	redisdao.Set(rt, payload, duration)

	return
}

// RefreshToken 拿取redis的payload，之后再单独调用createToken
func RefreshToken(rt string) (payload Payload) {
	p := redisdao.Get(rt)

	err := json.Unmarshal(p, &payload)
	if err != nil {
		log.Println("json unmarshal err : ", err)
	}

	return
}
