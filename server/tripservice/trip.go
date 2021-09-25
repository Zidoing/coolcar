package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

type Service struct {
	trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 30,
			FeeCent:     30,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			Status: trippb.TripStatus_PAID,
		},
	}, nil
}
