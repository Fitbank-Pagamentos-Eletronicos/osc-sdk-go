package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumerUnit(t *testing.T) {
	consumerUnit := domains.ConsumerUnit{
		Number: "123456789",
	}

	fmt.Println(consumerUnit.ToJson())

	assert.Equal(t, "123456789", consumerUnit.Number)
}
