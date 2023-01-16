package test

import (
	"osc-sdk-go/src/domains"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	identity := domains.Identity{
		Type:        domains.RG,
		Number:      "123456789",
		Issuer:      domains.SSP,
		State:       domains.BAHIA,
		IssuingDate: "01/01/2019",
	}

	assert.Equal(t, domains.RG, identity.Type)
	assert.Equal(t, "123456789", identity.Number)
	assert.Equal(t, domains.SSP, identity.Issuer)
	assert.Equal(t, domains.BAHIA, identity.State)
	assert.Equal(t, "01/01/2019", identity.IssuingDate)
}
