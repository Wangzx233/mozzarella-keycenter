package token

import (
	"log"
	"mozzarella-keycenter/api"
	"testing"
)

func TestVerifyToken(t *testing.T) {
	_, err := api.Key("123")
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	token, err := CreateToken("123", "111")
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
	err = VerifyToken(token)
	if err != nil {
		log.Println(err)
		t.Error(err)
		return
	}
}
