package main

import (
	"flag"
	"time"

	"github.com/micro/go-log"

	handler "github.com/akililab/balance/handler"
	subscriber "github.com/akililab/balance/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"

	balance "github.com/akililab/balance/proto/balance"
)

var (
	serverAddr string
)

func init() {
	flag.StringVar(&serverAddr, "server_address", "0.0.0.0:9080", "server address.")
	flag.Parse()
}

func main() {


	// New Service
	service := grpc.NewService(
		micro.Name("go.micro.srv.balance"),
		micro.RegisterTTL(time.Second*30),
		micro.Address(serverAddr),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	balance.RegisterBalanceServiceHandler(service.Server(), new(handler.Balance))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.balance", service.Server(), new(subscriber.Balance))

	// Register Function as Subscriber
	micro.RegisterSubscriber("topic.go.micro.srv.balance", service.Server(), subscriber.Handler)

	log.Logf("Service Run")

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
