package main

import (
	config "api_router/configs"
	"api_router/internal/service"
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
	"utils/etcd"
)

func main() {
	config.InitConfig()

	c := make(chan os.Signal, 1)
	go func() {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	}()

	etcdAddress := viper.GetString("etcd.address")
	serviceDiscovery, err := etcd.NewServiceDiscovery([]string{etcdAddress})
	if err != nil {
		log.Fatal(err)
	}
	defer serviceDiscovery.Close()

	err = serviceDiscovery.ServiceDiscovery("user_service")
	if err != nil {
		log.Fatal(err)
	}

	serviceAddr, _ := serviceDiscovery.GetService("user_service")
	conn, err := grpc.Dial(serviceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := service.NewUserServiceClient(conn)
	req := &service.UserRequest{
		Username: "zzz",
		Password: "123456",
	}

	resq, err := client.UserRegister(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", resq)
}
