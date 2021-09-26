package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	logger, err := newZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}

	server := grpc.NewServer()

	if err != nil {
		panic(err)
	}
	rentalpb.RegisterTripServiceServer(server, &trip.Service{
		Logger: logger,
	})
	fmt.Println("start serve")
	err = server.Serve(listen)
	logger.Fatal("cannot server", zap.Error(err))
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
