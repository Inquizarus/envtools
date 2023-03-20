package envtools

import (
	"encoding/json"
	"os"
)

// GetJSON retrieves a JSON string from an environment variable and then
// expands it with other environment variables available before unmarshaling
// it into the passed target that must be a pointer
func GetJSON(key string, target interface{}) (bool, error) {
	if data, found := os.LookupEnv(key); found {
		return true, json.Unmarshal([]byte(os.ExpandEnv(data)), target)
	}
	return false, nil
}
