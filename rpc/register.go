package rpc

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	rpc "mozzarella-book/rpc/pb"
	"net"
	"time"
)

const (
	serviceName = "mozzarella-book"
	ip          = "8.142.81.74"
	port        = 8082
	version     = "0.0.1"
)

func InitRegister() {

	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "175.24.203.115",
			Port:   8848,
		},
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		panic(err)
	}
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		Metadata:    map[string]string{"version": version, "up-time": time.Now().String()},
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	if success {
		//启动rpc server
		// 监听本地的8972端口
		lis, err := net.Listen("tcp", ":8082")
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
			return
		}
		s := grpc.NewServer()                                  // 创建gRPC服务器
		rpc.RegisterMozzarellaBookServer(s, &MozzarellaBook{}) // 在gRPC服务端注册服务

		reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
		// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
		// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
		err = s.Serve(lis)
		if err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	} else {
		log.Println(err)
	}
}
