package domains

//Struct to store the OAuth token
type AuthResponse struct {
	client_id     string
	client_secret string
	scopes        []string
}
