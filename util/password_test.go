package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	requires := require.New(t)
	password := RandomString(10)

	hashedPassword, err := HashPassword(password)
	requires.NoError(err)
	requires.NotEmpty(hashedPassword)

	err = ComparePassword(password, hashedPassword)
	requires.NoError(err)

	wrongPassword := RandomString(9)
	err = ComparePassword(wrongPassword, hashedPassword)
	requires.Error(err)

}
