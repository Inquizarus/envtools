package envtools_test

import (
	"os"
	"testing"

	"github.com/inquizarus/envtools"
	"github.com/stretchr/testify/assert"
)

func TestGetWithFallback(t *testing.T) {
	// Test with an existing environment variable.
	key := "MY_ENV_VAR"
	os.Setenv(key, "value")
	defer os.Unsetenv(key)

	// Test that the function returns the environment variable's value.
	value := envtools.GetWithFallback(key, "fallback")
	assert.Equal(t, "value", value)

	// Test with a non-existing environment variable.
	value = envtools.GetWithFallback("NON_EXISTING_VAR", "fallback")
	assert.Equal(t, "fallback", value)

	// Test with an empty string as fallback value.
	value = envtools.GetWithFallback("NON_EXISTING_VAR", "")
	assert.Equal(t, "", value)
}
