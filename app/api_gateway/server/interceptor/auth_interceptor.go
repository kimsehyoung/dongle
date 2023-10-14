package interceptor

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/kimsehyoung/dongle/app/auth/server/jwt"
	"google.golang.org/grpc"
)

// https://github.dev/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/auth/auth.go
// https://pkg.go.dev/github.com/grpc-ecosystem/go-grpc-middleware/v2#section-readme

var authSkipMethods = map[string]bool{
	"CreateAccount": true,
	"CreateToken":   true,
	"RefreshToken":  true,
}

func authFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	claim, err := jwt.VerifyJwt(token)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "login_id", claim.Subject)

	return ctx, nil
}

func authSkip(_ context.Context, c interceptors.CallMeta) bool {
	return !authSkipMethods[c.Method]
	// c.FullMethod() => '/dongle.speech.Speech/Recognize'
	// c.Service => 'dongle.speech.Speech'
	// c.Method => 'Recognize'
}

func AuthUnaryInterceptor() grpc.UnaryServerInterceptor {
	return selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(authFunc), selector.MatchFunc(authSkip))
}
