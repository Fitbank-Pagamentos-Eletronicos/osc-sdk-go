package test

import (
	"github.com/Fitbank-Pagamentos-Eletronicos/osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumerUnit(t *testing.T) {
	consumerUnit := domains.ConsumerUnit{
		Number: "123456789",
	}

	assert.Equal(t, "123456789", consumerUnit.Number)
}
