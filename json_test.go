package envtools_test

import (
	"os"
	"testing"

	"github.com/inquizarus/envtools"
	"github.com/stretchr/testify/assert"
)

func TestThatGetJSONWorks(t *testing.T) {
	type x struct {
		Name   string `json:"name"`
		Envvar string `json:"envvar"`
	}

	envVarKeyA := "TestThatGetJSONWorks_A"
	envVarKeyB := "TestThatGetJSONWorks_B"
	envVarValueA := `{"name":"foobar","envvar":"${` + envVarKeyB + `}"}`
	envVarValueB := "fizzbuzz"

	os.Setenv(envVarKeyA, envVarValueA)
	os.Setenv(envVarKeyB, envVarValueB)

	defer os.Unsetenv(envVarKeyA)
	defer os.Unsetenv(envVarKeyB)

	foo := x{}

	found, err := envtools.GetJSON(envVarKeyA, &foo)

	assert.True(t, found)
	assert.NoError(t, err)

	assert.Equal(t, "foobar", foo.Name)
	assert.Equal(t, "fizzbuzz", foo.Envvar)
}

func TestThatGetJSONReturnsFalseWhenNotFoundButWithoutError(t *testing.T) {
	found, err := envtools.GetJSON("TestThatGetJSONReturnsFalseWhenNotFoundButWithoutError", nil)
	assert.False(t, found)
	assert.NoError(t, err)
}

func TestThatGetJSONReturnsTrueAndErrorWhenBadJSONIsInEnvvar(t *testing.T) {
	key := "TestThatGetJSONReturnsTrueAndErrorWhenBadJSONIsInEnvvar"
	value := `{"foo":",}`

	os.Setenv(key, value)
	defer os.Unsetenv(key)

	found, err := envtools.GetJSON(key, &struct{}{})
	assert.True(t, found)
	assert.Error(t, err)
}
