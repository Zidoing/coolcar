package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	trip "coolcar/tripservice"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {

	go startGRPCGateway()

	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	server := grpc.NewServer()
	trippb.RegisterTripServiceServer(server, &trip.Service{})
	log.Fatal(server.Serve(listen))
}

func startGRPCGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(ctx, mux, ":8081", []grpc.DialOption{
		grpc.WithInsecure(),
	})

	if err != nil {
		log.Fatalf("cannot start grpc gateway %v", err)
	}
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("cannot listen and server %v", err)
	}
}
