package util

import "encoding/json"

func PrettyJSON(v interface{}) string {
	b, _ := json.MarshalIndent(v, " ", " ")
	return string(b)
}
