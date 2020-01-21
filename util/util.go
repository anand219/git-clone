package util

import (
	"encoding/base64"
	"encoding/json"
)

func PrettyJSON(v interface{}) string {
	b, _ := json.MarshalIndent(v, " ", " ")
	return string(b)
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
