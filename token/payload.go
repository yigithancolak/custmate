package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	OrganizationID string `json:"organization_id"`
	jwt.StandardClaims
}

func NewPayload(id string, duration time.Duration) (*Payload, error) {

	payload := &Payload{
		OrganizationID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}

	return payload, nil
}

// func (payload *Payload) Valid() error {
// 	if time.Now().After(payload.ExpiredAt) {
// 		return ErrExpiredToken
// 	}
// 	return nil
// }
