package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	*rentalpb.UnimplementedTripServiceServer
	Logger *zap.Logger
	Mongo  *dao.Mongo
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		panic(err)
	}
	s.Logger.Info("create trip", zap.String("start", req.Start.String()), zap.String("account_id", aid.String()))
	trip := &rentalpb.Trip{
		AccountId: aid.String(),
		CarId:     "",
		Start: &rentalpb.LocationStatus{
			Location: req.Start,
			FeeCent:  0,
			KmDriven: 0,
			PoiName:  "",
		},
		Current: nil,
		End:     nil,
		Status:  rentalpb.TripStatus_FINISHED,
	}
	s.Mongo.CreateTrip(c, trip)

	return &rentalpb.TripEntity{
		Id:   "222",
		Trip: trip,
	}, nil

	return nil, status.Error(codes.Unimplemented, "")
}
