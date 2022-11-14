package requests

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"modulo/src/domains"
	jsonUtil "modulo/src/utils/json"
	"modulo/src/utils/jwt"
	"net/http"
	"strings"
	"errors"
)

type AuthResponse struct {
	Access_token string `json:"access_token"`
	Expire_at    string `json:"expire_at"`
}

type AuthorizationRequest struct {
	Client_id string `json:"client_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type AuthBodyRequest struct {
	Scopes []string `json:"scopes"`
}

var userDataBase = domains.Auth{
	Client_id:     "104",
	Client_secret: "4dbe3aa7-8ce9-43a4-9298-73b700e712bb:1b364af124250aa09461f33161c3d96e551d822080fe1bd977aa66d7ec9378c8",
	Scopes:        []string{"api-external"},
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func Auth(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	var user, err = getUserByCredentiais(request.Header)

	
	if err != nil{
		fmt.Fprintln(writer, jsonUtil.Encode(ErrorResponse{Message: err.Error()}))
		
		return
	}

	reqBody, _ := ioutil.ReadAll(request.Body)
	var post AuthBodyRequest
	jsonUtil.Decode(reqBody, &post)

	var token = jwt.GenerateToken(user.Client_id, user.Username, post.Scopes)

	var authResponse = AuthResponse{
		Access_token: token.Access_token,
		Expire_at:    token.Expire_at,
	}
	

	fmt.Println("authResponse", authResponse)
	
	fmt.Fprintln(writer, jsonUtil.Encode(authResponse))

}

func getUserByCredentiais(header http.Header) (AuthorizationRequest, error) {

	var response = header.Get("Authorization")

	if response == "" {
		return AuthorizationRequest{}, errors.New("Internal Server Error")
	}
	

	sUserB64, _ := base64.StdEncoding.DecodeString(removeBasic(response))

	if string(sUserB64) != userDataBase.Client_secret {
		
		return AuthorizationRequest{}, errors.New("User not found")
	}

	return userFromBase64(string(sUserB64)), nil

}

func removeBasic (response string) string {
	return strings.Replace(response, "Basic ", "", 1)
}


func userFromBase64(userB64 string) AuthorizationRequest {

	var decodedUser = strings.Split(userB64, ":")

	return AuthorizationRequest{
		Client_id: userDataBase.Client_id,
		Username:  decodedUser[0],
		Password:  decodedUser[1],
	}
}
