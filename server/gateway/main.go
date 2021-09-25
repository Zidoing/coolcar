package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

func main() {
	c := context.Background()
	c, cancelFunc := context.WithCancel(c)
	defer cancelFunc()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{},
			UnmarshalOptions: protojson.UnmarshalOptions{},
		},
	))

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		c, mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatalf("cannot register auth service :%v", err)
	}

	log.Fatal(http.ListenAndServe(":8080", mux))

}
