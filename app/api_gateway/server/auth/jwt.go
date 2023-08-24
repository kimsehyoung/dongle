package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	auth_interceptor "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/kimsehyoung/dongle/api/proto/gen/go/authpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenType string

const (
	Token_ACCESS  TokenType = "access"
	Token_REFRESH TokenType = "refresh"
)

const (
	// TODO: This will be replaced to 'Valut'.
	secretKey            = "temporary_secret_key"
	accessTokenDuration  = 30 * time.Minute
	refreshTokenDuration = 30 * 24 * time.Hour
	refreshTokenReissue  = 10 * 24 * time.Hour
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	Role authpb.Role `json:"role"`
}

var authSkipMethods = map[string]bool{
	"CreateAccount": true,
	"CreateToken":   true,
	"RefreshToken":  true,
}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := auth_interceptor.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	_, err = VerifyJwt(token)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func AuthSkip(_ context.Context, c interceptors.CallMeta) bool {
	return !authSkipMethods[c.Method]
	// c.FullMethod() => '/dongle.speech.Speech/Recognize'
	// c.Service => 'dongle.speech.Speech'
	// c.Method => 'Recognize'
}

func CreateJwt(loginId string, role authpb.Role, tokenType TokenType) (string, error) {
	duration := accessTokenDuration
	if tokenType == Token_REFRESH {
		duration = refreshTokenDuration
	}

	claims := &MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sehyoung",
			Subject:   "dongle",
			Audience:  []string{loginId},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			// NotBefore:
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// ID:
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyJwt(token string) (*MyCustomClaims, error) {

	toeknRes, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err == nil {
		if claims, ok := toeknRes.Claims.(*MyCustomClaims); ok && toeknRes.Valid {
			return claims, nil
		}
	}

	if err == jwt.ErrTokenMalformed {
		return nil, status.Errorf(codes.InvalidArgument, "That's not even a token: %v", err)
	} else if err == jwt.ErrTokenExpired {
		return nil, status.Errorf(codes.Unauthenticated, "Token is expired: %v", err)
	} else if err == jwt.ErrTokenNotValidYet {
		return nil, status.Errorf(codes.Unauthenticated, "Token is not active yet: %v", err)
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Couldn't handle this token: %v", err)
	}
}
