package envtools_test

import (
	"os"
	"testing"

	"github.com/inquizarus/envtools"
	"github.com/stretchr/testify/assert"
)

func TestGetRequiredOrPanic(t *testing.T) {
	// Test case 1: key exists in environment
	os.Setenv("KEY1", "value1")
	defer os.Unsetenv("KEY1")

	assert.Equal(t, envtools.GetRequiredOrPanic("KEY1"), "value1")

	// Test case 2: key does not exist in environment
	assert.PanicsWithError(t, "environment variable with key KEY2 is not set", func() {
		envtools.GetRequiredOrPanic("KEY2")
	})
}

func TestGetRequiredOrError(t *testing.T) {
	// Test case 1: key exists in environment
	os.Setenv("KEY1", "value1")
	defer os.Unsetenv("KEY1")

	value, err := envtools.GetRequiredOrError("KEY1")
	assert.Equal(t, value, "value1")
	assert.NoError(t, err)

	// Test case 2: key does not exist in environment
	value, err = envtools.GetRequiredOrError("KEY2")
	assert.Equal(t, value, "")
	assert.EqualError(t, err, "environment variable with key KEY2 is not set")
}
