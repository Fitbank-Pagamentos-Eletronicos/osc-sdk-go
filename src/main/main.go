package main

import (
    "regexp"
    "fmt"
)

func Normalize(src string) string {
    var re = regexp.MustCompile("[^a-zA-Z0-9]+")
    return re.ReplaceAllString(src, "")
}

func main() {
   var res = Normalize("abc&%$#@123")
    fmt.Println(res)
}
