package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenType string

const (
	ACCESS_TOKEN  TokenType = "access"
	REFRESH_TOKEN TokenType = "refresh"
)

const (
	// TODO: This will be replaced to 'Valut'.
	secretKey            = "temporary_secret_key"
	accessTokenDuration  = 30 * time.Minute
	RefreshTokenDuration = 1 * 24 * time.Hour
)

type MyCustomClaims struct {
	jwt.RegisteredClaims
	RoleId int `json:"roleId"`
}

func CreateJwt(roleId int, loginId string, tokenType TokenType) (string, error) {
	duration := accessTokenDuration
	if tokenType == REFRESH_TOKEN {
		duration = RefreshTokenDuration
	}

	claims := &MyCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sehyoung",
			Subject:   loginId,
			Audience:  []string{"dognle"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			// NotBefore:
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// ID:
		},
		RoleId: roleId,
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
