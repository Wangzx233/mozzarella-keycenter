package main

import (
	"mozzarella-keycenter/redisdao"
	"mozzarella-keycenter/rpc"
)

func main() {
	redisdao.InitRedis()
	rpc.InitRegister()
}
