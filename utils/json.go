package utils

import "encoding/json"

func ToJson(v interface{}) (string, error) {
	str, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
