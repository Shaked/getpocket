package auth

type Authenticated interface {
	AccessToken() string
}

type User interface {
	Username() string
}

type PocketUser struct {
	PAccessToken string `json:"access_token"`
	PUsername    string `json:"username"`
}

func NewUser() *PocketUser {
	return &PocketUser{}
}

func (u *PocketUser) SetAccessToken(accessToken string) *PocketUser {
	u.PAccessToken = accessToken
	return u
}

func (u *PocketUser) SetUsername(username string) *PocketUser {
	u.PUsername = username
	return u
}

func (u *PocketUser) AccessToken() string {
	return u.PAccessToken
}

func (u *PocketUser) Username() string {
	return u.PUsername
}
