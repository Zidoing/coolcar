package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Service struct {
	*authpb.UnimplementedAuthServiceServer
	TokenGenerator TokenGenerator
	TokenExpire    time.Duration
	Logger         *zap.Logger
	Mongo          *dao.Mongo
}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code", zap.String("code", req.Code))

	accountID, err := s.Mongo.ResolveAccountID(c, req.Code)

	token, err := s.TokenGenerator.GenerateToken(accountID, s.TokenExpire)
	if err != nil {
		s.Logger.Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	if err != nil {
		s.Logger.Error("cannot resolve account id")
		return nil, status.Error(codes.Internal, "")
	}
	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TokenExpire.Seconds()) + 1,
	}, nil
}
