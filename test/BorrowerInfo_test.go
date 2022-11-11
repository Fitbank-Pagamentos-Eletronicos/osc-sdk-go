package test

import (
	"fmt"
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorrowerInfo(t *testing.T) {
	borrowerInfo := domains.BorrowerInfo{
		IsRegistered: true,
		IsBlocked:    true,
	}

	fmt.Println(borrowerInfo.ToJson())

	assert.Equal(t, true, borrowerInfo.IsRegistered)
	assert.Equal(t, true, borrowerInfo.IsBlocked)

}
