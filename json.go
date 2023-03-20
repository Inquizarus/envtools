package envtools

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// UnmarshalJSON retrieves a JSON string from an environment variable and then
// expands it with other environment variables available before unmarshaling
// it into the passed target that must be a pointer
func UnmarshalJSON(key string, target interface{}) (bool, error) {
	if data, found := os.LookupEnv(key); found {
		return true, json.Unmarshal([]byte(os.ExpandEnv(data)), target)
	}
	return false, nil
}

func GetJSONData(key string) (io.Reader, bool) {
	if data, found := os.LookupEnv(key); found {
		return bytes.NewBuffer([]byte(os.ExpandEnv(data))), found
	}
	return nil, false
}
