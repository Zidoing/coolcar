package server

import (
	"coolcar/shared/auth"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	Logger            *zap.Logger
	RegisterFunc      func(server *grpc.Server)
}

func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)

	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var opts []grpc.ServerOption

	if c.AuthPublicKeyFile != "" {
		interceptor, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			panic(err)
		}
		opts = append(opts, grpc.UnaryInterceptor(interceptor))
	}
	server := grpc.NewServer(opts...)

	if err != nil {
		panic(err)
	}
	c.RegisterFunc(server)

	fmt.Println("start serve ", c.Name)
	return server.Serve(listen)
}
