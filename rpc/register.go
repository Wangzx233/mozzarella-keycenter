package rpc

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"time"
)

const (
	serviceName = "mozzarella-keycenter"
	ip          = "8.142.81.74"
	port        = 8901
	version     = "0.0.1"
)

func InitRegister() {

	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "error",
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
		InitRpc()
	} else {
		log.Println(err)
	}
}
