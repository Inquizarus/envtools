package envtools_test

import (
	"io"
	"os"
	"testing"

	"github.com/inquizarus/envtools"
	"github.com/stretchr/testify/assert"
)

func TestThatUnmarshalJSONWorks(t *testing.T) {
	type x struct {
		Name   string `json:"name"`
		Envvar string `json:"envvar"`
	}

	keyA := "TestThatUnmarshalJSONWorks_A"
	keyB := "TestThatUnmarshalJSONWorks_B"
	valueA := `{"name":"foobar","envvar":"${` + keyB + `}"}`
	valueB := "fizzbuzz"

	os.Setenv(keyA, valueA)
	os.Setenv(keyB, valueB)

	defer os.Unsetenv(keyA)
	defer os.Unsetenv(keyB)

	foo := x{}

	found, err := envtools.UnmarshalJSON(keyA, &foo)

	assert.True(t, found)
	assert.NoError(t, err)

	assert.Equal(t, "foobar", foo.Name)
	assert.Equal(t, "fizzbuzz", foo.Envvar)
}

func TestThatUnmarshalJSONReturnsFalseWhenNotFoundButWithoutError(t *testing.T) {
	found, err := envtools.UnmarshalJSON("TestThatUnmarshalJSONReturnsFalseWhenNotFoundButWithoutError", nil)
	assert.False(t, found)
	assert.NoError(t, err)
}

func TestThatUnmarshalJSONReturnsTrueAndErrorWhenBadJSONIsInEnvvar(t *testing.T) {
	key := "TestThatUnmarshalJSONReturnsTrueAndErrorWhenBadJSONIsInEnvvar"
	value := `{"foo":",}`

	os.Setenv(key, value)
	defer os.Unsetenv(key)

	found, err := envtools.UnmarshalJSON(key, &struct{}{})
	assert.True(t, found)
	assert.Error(t, err)
}

func TestThatGetJSONDataWorks(t *testing.T) {
	keyA := "TestThatGetJSONDataWorksA"
	keyB := "TestThatGetJSONDataWorksB"
	valueA := `{"envvar":"${` + keyB + `}"}`
	valueB := "foobar"

	os.Setenv(keyA, valueA)
	os.Setenv(keyB, valueB)

	defer os.Unsetenv(keyA)
	defer os.Unsetenv(keyB)

	r, found := envtools.GetJSONData(keyA)

	assert.NotNil(t, r)
	assert.True(t, found)

	data, err := io.ReadAll(r)

	assert.NoError(t, err)

	assert.Equal(t, `{"envvar":"`+valueB+`"}`, string(data))

}
