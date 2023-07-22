package server

import (
	"context"
	"fmt"

	"github.com/inveracity/svelte-grpc-stream/internal/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authMgr *auth.Auth
}

func NewAuthInterceptor(authMgr *auth.Auth) *AuthInterceptor {
	return &AuthInterceptor{
		authMgr: authMgr,
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("no metadata found")
	}

	token := md.Get("jwt")

	authed, err := interceptor.authMgr.VerifyUserToken(token[0])
	if err != nil || !authed {
		return fmt.Errorf("user not authorized")
	}
	return nil
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		err := interceptor.authorize(stream.Context())
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		err := interceptor.authorize(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
