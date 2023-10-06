package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/yigithancolak/custmate/util"
)

func TestJWTMaker(t *testing.T) {
	asserts := assert.New(t)

	config, err := util.LoadConfig("..", "test", "env")
	asserts.NoError(err)

	maker, err := NewJWTMaker(config)
	asserts.NoError(err)

	t.Run("Create and Verify Token", func(t *testing.T) {
		testID := uuid.New().String()
		token, payload, err := maker.CreateToken(testID, config.AccessTokenDuration)
		asserts.NoError(err)
		asserts.NotEmpty(token)
		asserts.Equal(testID, payload.OrganizationID)

		// Verify the token
		verifiedPayload, err := maker.VerifyToken(token)
		asserts.NoError(err)
		asserts.Equal(testID, verifiedPayload.OrganizationID)
	})

	t.Run("Invalid Token", func(t *testing.T) {
		token := "invalid.token.string"
		_, err := maker.VerifyToken(token)
		asserts.Error(err)
		asserts.Equal(ErrInvalidToken, err)
	})

	t.Run("Expired Token", func(t *testing.T) {
		// Create a token that's already expired
		testID := uuid.New().String()
		token, _, err := maker.CreateToken(testID, -24*time.Hour)
		asserts.NoError(err)
		asserts.NotEmpty(token)

		_, err = maker.VerifyToken(token)
		asserts.Error(err)
		asserts.Equal(ErrExpiredToken, err)
	})
}
