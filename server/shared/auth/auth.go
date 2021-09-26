package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"os"
	"strings"
)

type tokenVerifier interface {
	Verify(token string) (string, error)
}

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(publicKeyFile)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		panic(err)
	}
	i := &interceptor{
		verify: &token.JWTTokenVerifier{PublicKey: publicKey},
	}
	return i.HandleReq, nil
}

type interceptor struct {
	verify tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		panic(err)
	}

	aid, err := i.verify.Verify(tkn)
	if err != nil {
		panic(err)
	}
	return handler(ContextWithAccountID(ctx, aid), req)

}

func tokenFromContext(c context.Context) (string, error) {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		panic(ok)
	}
	tkn := ""
	for _, v := range m["authorization"] {
		if strings.HasPrefix(v, "Bearer") {
			tkn = v[len("Bearer "):]
		}
	}
	if tkn == "" {
		return "", nil
	}
	return tkn, nil
}

type accountIDKey struct {
}

func ContextWithAccountID(c context.Context, aid string) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func AccountIDFromContext(c context.Context) (string, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(string)
	if !ok {
		return "", nil
	}
	return aid, nil

}
