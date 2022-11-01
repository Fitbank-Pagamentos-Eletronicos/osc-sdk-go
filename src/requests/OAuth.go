package requests

//Struct to store the OAuth token
type AuthResponse struct {
	client_id     string
	client_secret string
	scopes        []string
}

//Function to create a new OAuth token
func OAuth(client_id string, client_secret string, scopes []string) AuthResponse {
	return AuthResponse{client_id, client_secret, scopes}
}
