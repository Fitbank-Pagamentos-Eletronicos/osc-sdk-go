package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPipeline(t *testing.T) {
	pipeline := domains.Pipeline{
		Id:          "1",
		Status:      domains.SIGNUP_ANALISIS,
		Cpf:         "123456789",
		Name:        "Joao",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
	}

	assert.Equal(t, pipeline.Id, "1")
	assert.Equal(t, pipeline.Status, domains.SIGNUP_ANALISIS)
	assert.Equal(t, pipeline.Cpf, "123456789")
	assert.Equal(t, pipeline.Name, "Joao")
	assert.Equal(t, pipeline.DateCreated, "2020-01-01")
	assert.Equal(t, pipeline.LastUpdated, "2020-01-01")
}
