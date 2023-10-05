package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	requires := require.New(t)
	existing := RandomString(6)
	notExisting := RandomString(4)
	var randomStringSlice = []string{RandomString(4), RandomString(6), existing}

	ok := Contains[string](randomStringSlice, existing)
	requires.True(ok)

	ok = Contains[string](randomStringSlice, notExisting)
	requires.False(ok)

}
