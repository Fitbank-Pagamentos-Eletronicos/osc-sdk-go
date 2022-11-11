package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelineProposal(t *testing.T) {
	pipelineProposal := domains.PipelineProposal{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "123456789",
		Name:        "Alexandre",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
	}
	
	fmt.Println(pipelineProposal.ToJson())

	assert.Equal(t, pipelineProposal.Id, "1")
	assert.Equal(t, pipelineProposal.Status, "Aprovado")
	assert.Equal(t, pipelineProposal.Cpf, "123456789")
	assert.Equal(t, pipelineProposal.Name, "Alexandre")
	assert.Equal(t, pipelineProposal.DateCreated, "2020-01-01")
	assert.Equal(t, pipelineProposal.LastUpdated, "2020-01-01")
}
