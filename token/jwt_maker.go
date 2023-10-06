package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/yigithancolak/custmate/util"
)

type JWTMaker struct {
	secretKey            string
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

const minSecretKeySize = 32

func NewJWTMaker(config *util.Config) (*JWTMaker, error) {
	if len(config.TokenSymmetricKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size. must be at leat %d characters", minSecretKeySize)
	}

	return &JWTMaker{
		secretKey:            config.TokenSymmetricKey,
		AccessTokenDuration:  config.AccessTokenDuration,
		RefreshTokenDuration: config.RefreshTokenDuration,
	}, nil
}

func (maker JWTMaker) CreateToken(id string, duration time.Duration) (string, *Payload, error) {

	payload, err := NewPayload(id, duration)

	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok && ve.Errors == jwt.ValidationErrorExpired {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	if payload.ExpiresAt < time.Now().Local().Unix() {
		return nil, ErrExpiredToken
	}

	return payload, nil
}
