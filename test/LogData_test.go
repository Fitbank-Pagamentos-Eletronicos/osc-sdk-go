// Teste logData
package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogData(t *testing.T) {
	logData := domains.LogData{
		Latitude:      -16.6982283,
		Longitude:     -49.2581201,
		OccurenceData: "2019-08-21T14:31:17.459Z",
		UserAgent:     "Mozilla/5.0",
		Ip:            "0.0.0.0",
		Mac:           "00:00:00:00:00:00",
	}

	fmt.Println(logData.ToJson())

	assert.Equal(t, -16.6982283, logData.Latitude)
	assert.Equal(t, -49.2581201, logData.Longitude)
	assert.Equal(t, "2019-08-21T14:31:17.459Z", logData.OccurenceData)
	assert.Equal(t, "Mozilla/5.0", logData.UserAgent)
	assert.Equal(t, "0.0.0.0", logData.Ip)
	assert.Equal(t, "00:00:00:00:00:00", logData.Mac)

}
