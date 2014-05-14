package auth

type AuthUser struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}
