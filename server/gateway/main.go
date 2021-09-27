package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
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

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{name: "auth", addr: "localhost:8081", registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint},
		{name: "rental", addr: "localhost:8081", registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint},
	}

	for _, s := range serverConfig {
		err := s.registerFunc(c, mux, s.addr, []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			log.Fatalf("cannot register %s service :%v", s.name, err)
		}
	}

	log.Fatal(http.ListenAndServe(":8080", mux))

}
