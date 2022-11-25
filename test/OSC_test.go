package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "modulo/src/requests"
    "modulo/src/domains"

)

func TestOSC(t *testing.T) {

        getToken := domains.AuthSucess{
            Access_token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
            Expire_at: "2022-11-26T10:56:24.038Z",
        }

       if getToken.Access_token == "" || getToken.Expire_at == "" {
            assert.Equal(t, getToken, requests.GetToken())
            assert.Equal(t, getToken, requests.GetToken())
       }

        normalized := requests.Normalize
        assert.Equal(t, normalized("@Pé de moleque"), "-pe-de-moleque")


}


