package main

import (
	"modulo/src/requests"
	"net/http"
)

func main() {

	http.HandleFunc("/client/auth", requests.Auth)

	http.ListenAndServe(":8080", nil)
}
