package main

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"modulo/src/domains"
	"modulo/src/requests"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var auth domains.AuthSucess

func GetToken() string {
    if auth.Access_token == "" || auth.Expire_at == "" {
		auth = Auth()
	   	return auth.Access_token
	}

	var expireAt, _ = time.Parse("2006-01-02T15:04:05.000Z", auth.Expire_at)

	if time.Now().After(expireAt) {
		auth = Auth()
		return auth.Access_token
	}

	return auth.Access_token
}

func Normalize(name string) string {
	toConvert := regexp.MustCompile("[^A-z0-9]+")
	return toConvert.ReplaceAllString(strings.ToLower(strings.Trim(RemoveAccents(name), " ")), "-")
}

func RemoveAccents(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

func Auth() domains.AuthSucess {
	return requests.OAuth()
}

func main() {
	GetToken()
	fmt.Println(Normalize("João da Silva"))
}
