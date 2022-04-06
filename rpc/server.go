package rpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mozzarella-keycenter/key"
	"mozzarella-keycenter/pb"
	"mozzarella-keycenter/token"
	"net"
)

type KeyCenter struct {
}

func (k *KeyCenter) CreateToken(c context.Context, req *pb.CreateTokenReq) (res *pb.CreateTokenResp, err error) {
	ac, rt, err := token.CreateToken(req.Domain, req.Uid)
	if err != nil {
		log.Println(err)
		return
	}
	return &pb.CreateTokenResp{
		Token:        ac,
		RefreshToken: rt,
	}, nil
}

func (k *KeyCenter) VerifyToken(c context.Context, req *pb.VerifyTokenReq) (res *pb.VerifyTokenResp, err error) {
	err = token.VerifyToken(req.Token)
	if err != nil {
		err = errors.New("verify err")
		return &pb.VerifyTokenResp{Ok: false}, err
	}
	return &pb.VerifyTokenResp{Ok: true}, nil
}

func (k *KeyCenter) Key(c context.Context, req *pb.KeyRequest) (res *pb.KeyReturn, err error) {
	Key, err := key.GetKey(req.Domain)
	if err != nil {
		return
	}
	return &pb.KeyReturn{PublicKey: Key.PublicKey}, nil
}

func (k *KeyCenter) Ping(c context.Context, req *pb.PingRequest) (res *pb.PingReply, err error) {
	return &pb.PingReply{Pong: "pong"}, nil
}

func (k *KeyCenter) RefreshToken(c context.Context, req *pb.RefreshTokenReq) (res *pb.RefreshTokenResp, err error) {
	payload, err := token.RefreshToken(req.GetRt())
	if err != nil {
		return
	}
	ac, rt, err := token.CreateToken(payload.Subject, payload.Uid)
	if err != nil {
		return
	}

	return &pb.RefreshTokenResp{
		Token:        ac,
		RefreshToken: rt,
	}, nil
}

func InitRpc() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8901")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器

	pb.RegisterMozzarellaBookServer(s, &KeyCenter{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
