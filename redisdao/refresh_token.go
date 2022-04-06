package redisdao

import (
	"log"
	"time"
)

// Set 存储rt
func Set(key string, value []byte, duration time.Duration) {
	err := rdb.Set(key, value, duration).Err()
	if err != nil {
		log.Println("redis set err : ", err)
	}
	return
}

func Get(key string) (value []byte) {
	value, err := rdb.Get(key).Bytes()
	if err != nil {
		log.Println("redis set err : ", err)
	}
	return
}
