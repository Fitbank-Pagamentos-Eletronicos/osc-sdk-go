package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPipelineExpire(t *testing.T){
	pipelineExpire := domains.PipelineExpire{
		Id:          "1",
		Status:      "Aprovado",
		Cpf:         "123456789",
		Name:        "Alexandre",
		DateCreated: "2020-01-01",
		LastUpdated: "2020-01-01",
	}

	fmt.Println(pipelineExpire.ToJson())

	assert.Equal(t, pipelineExpire.Id, "1")
	assert.Equal(t, pipelineExpire.Status, "Aprovado")
	assert.Equal(t, pipelineExpire.Cpf, "123456789")
	assert.Equal(t, pipelineExpire.Name, "Alexandre")
	assert.Equal(t, pipelineExpire.DateCreated, "2020-01-01")
}