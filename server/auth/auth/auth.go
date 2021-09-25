package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	*authpb.UnimplementedAuthServiceServer
	Logger *zap.Logger
	Mongo  *dao.Mongo
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))

	accountID, err := s.Mongo.ResolveAccountID(c, req.Code)

	if err != nil {
		s.Logger.Error("cannot resolve account id")
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		AccessToken: "token for " + accountID,
		ExpiresIn:   111,
	}, nil
}
