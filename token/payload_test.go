package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPayload(t *testing.T) {
	// Mock the current time
	now := time.Now().UTC()
	duration := 24 * time.Hour
	testID := uuid.New().String()

	// Use the mocked time in your NewPayload function
	payload, err := NewPayload(testID, duration)

	assert.NoError(t, err)
	assert.Equal(t, testID, payload.OrganizationID)
	assert.Equal(t, now.Add(duration).Unix(), payload.ExpiresAt)
}
