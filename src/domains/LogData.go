package domains

import json2 "encoding/json"

type LogData struct {
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	OccurrenceDate string  `json:"occurrenceDate"`
	UserAgent      string  `json:"userAgent"`
	Ip             string  `json:"ip"`
	Mac            string  `json:"mac"`
}

func (a *LogData) ToJson() string {
	json, _ := json2.Marshal(a)
	return string(json)
}
