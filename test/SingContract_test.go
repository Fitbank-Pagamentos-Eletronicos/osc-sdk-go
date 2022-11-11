package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigContract(t *testing.T) {
	singContract := domains.SingContract{
		AceptedCheckSum: "123456789",
	}

	fmt.Println(singContract.ToJson())

	assert.Equal(t, "123456789", singContract.AceptedCheckSum)
}
