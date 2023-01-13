package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigContract(t *testing.T) {
	singContract := domains.SingContract{
		AcceptedCheckSum: "123456789",
	}

	assert.Equal(t, "123456789", singContract.AcceptedCheckSum)
}
