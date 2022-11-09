package test

import (
	"modulo/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBorrowerInfo(t *testing.T) {
	borrowerInfo := domains.BorrowerInfo{
		IsRegistered: true,
		IsBlocked:    true,
	}

	assert.Equal(t, true, borrowerInfo.IsRegistered)
	assert.Equal(t, true, borrowerInfo.IsBlocked)

}
