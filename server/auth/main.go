package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"time"
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

	pkFile, err := os.Open("auth/private.key")
	if err != nil {
		panic(err)
	}

	pkBytes, err := io.ReadAll(pkFile)

	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)

	if err != nil {
		panic(err)
	}

	authpb.RegisterAuthServiceServer(server, &auth.Service{
		Logger:         logger,
		Mongo:          dao.NewMongo(mc.Database("coolcar")),
		TokenExpire:    time.Minute * 100,
		TokenGenerator: token.NewJWTTokenGen("coolcar/auth", key),
	})

	err = server.Serve(listen)

	logger.Fatal("cannot server", zap.Error(err))
}

func newZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
