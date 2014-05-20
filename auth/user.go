package auth

type User struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}
