package domains

import "encoding/json"

type LogData struct {
	Latitude      float64
	Longitude     float64
	OccurenceData string
	UserAgent     string
	Ip            string
	Mac           string
}

func (a *LogData) ToJson() string {
	json, _ := json.Marshal(a)
	return string(json)
}
