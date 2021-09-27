package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create logger:%v", err)
	}

	server.RunGRPCServer(&server.GRPCConfig{
		Name:              "rental",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/token/public.key",
		Logger:            logger,
		RegisterFunc: func(server *grpc.Server) {
			rentalpb.RegisterTripServiceServer(server, &trip.Service{
				Logger: logger,
			})
		},
	})

}
