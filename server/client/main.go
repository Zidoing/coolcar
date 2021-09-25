package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server:%v", err)
	}

	tripServiceClient := trippb.NewTripServiceClient(conn)
	response, err := tripServiceClient.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip445"})
	if err != nil {
		log.Fatalf("cannot call Gettrip:%v", err)
	}
	fmt.Println(response)
}
