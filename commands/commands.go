package commands

import (
	"../auth"
	"../utils"
)

var (
	URLs = map[string]string{
		"Add": "https://getpocket.com/v3/add",
	}
)

type Response interface{}
type Executable interface {
	Exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error)
}

type Commands struct {
	user        *auth.User
	consumerKey string
	request     utils.HttpRequest
}

func New(user *auth.User, consumerKey string) *Commands {
	request := utils.NewRequest()
	return &Commands{
		user:        user,
		request:     request,
		consumerKey: consumerKey,
	}
}

func (c *Commands) Exec(command Executable) (Response, error) {
	resp, err := command.Exec(c.user, c.consumerKey, c.request)
	return resp, err
}

type Item struct {
	Id        string `json:"item_id"`
	NormalURL string `json:"normal_url"`
}
