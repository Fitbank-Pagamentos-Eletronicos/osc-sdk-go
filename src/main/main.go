package main

import (
	"fmt"
	"modulo/src/requests"
)

func main() {
    requests.GetToken()
    normalized := requests.Normalize
    fmt.Println("========STRING NORMALIZADA=========")
    fmt.Println(normalized(" @Pé de moleque "))


}
