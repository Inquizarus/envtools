package envtools

import (
	"errors"
	"fmt"
	"os"
)

func GetRequiredOrPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(errors.New("environment variable with key " + key + " is not set"))
}

func GetRequiredOrError(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", fmt.Errorf("environment variable with key %s is not set", key)
}
