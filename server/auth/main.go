package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}
	ctx := context.Background()
	mc, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.50.5:27017/?readPreference=primary&appname=mongodb-vscode%200.6.10&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect mongodb", zap.Error(err))
	}

	server := grpc.NewServer()
	authpb.RegisterAuthServiceServer(server, &auth.Service{
		Logger: logger,
		Mongo:  dao.NewMongo(mc.Database("coolcar")),
	})

	err = server.Serve(listen)

	logger.Fatal("cannot server", zap.Error(err))
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
