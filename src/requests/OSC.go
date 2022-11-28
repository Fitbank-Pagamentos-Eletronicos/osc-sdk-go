package requests

import (
    "modulo/src/domains"
    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
    "regexp"
    "time"
    "strings"
    "unicode"


)

var auth domains.AuthSucess

func GetToken () string {
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

func Normalize (name string) string {
    toConvert := regexp.MustCompile("[^A-z0-9]+")
    return toConvert.ReplaceAllString(strings.ToLower(strings.Trim(RemoveAccents(name), " ")), "-")
}

func RemoveAccents (str string) string {
    format := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
    result, _, _ := transform.String(format, str)
    return result
}

func isMn (runner rune) bool {
    return unicode.Is(unicode.Mn, runner)
}

func Auth () domains.AuthSucess {
    return OAuth()
}

