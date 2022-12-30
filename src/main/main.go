package main

import (
    "modulo/src/requests"
    "fmt"
)

func main() {

    fmt.Println("==================Requisição de CustomerServerNumber==================")
    requests.CustomerServiceNumberGET()

}
