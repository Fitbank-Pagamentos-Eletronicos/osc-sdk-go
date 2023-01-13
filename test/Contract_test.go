package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContract(t *testing.T) {
	logData := domains.LogData{
		Latitude:       -16.6982283,
		Longitude:      -49.2581201,
		OccurrenceDate: "2019-08-21T14:31:17.459Z",
		UserAgent:      "Mozilla/5.0",
		Ip:             "0.0.0.0",
		Mac:            "00:00:00:00:00:00",
	}

	fmt.Println(logData.ToJson())

	contract := domains.Contract{
		AcceptedCheckSum: []string{"1234567890"},
		LogData:          logData,
	}

	fmt.Println(contract.ToJson())

	assert.Equal(t, "acceptedCheckSum", contract.AcceptedCheckSum)
	assert.Equal(t, logData, contract.LogData)
}
