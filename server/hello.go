package main

import (
	trippb "coolcar/proto/gen/go"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start:       "abc",
		End:         "def",
		DurationSec: 30,
		FeeCent:     30,
	}
	fmt.Println(&trip)
	b, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%X\n", b)

	jsonTrip, err := json.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("jsonTrip %s\n", jsonTrip)

	var trip2 trippb.Trip
	err = proto.Unmarshal(b, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)
}
