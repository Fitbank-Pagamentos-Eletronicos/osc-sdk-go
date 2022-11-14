package domains

import (
	"encoding/json"
)

func Encode(data any) string {
	json, _ := json.Marshal(data)
	return string(json)
}

func Decode(data []byte, v any) {
	json.Unmarshal(data, &v)

}
