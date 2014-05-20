package commands

import (
	"log"
	"net/url"

	"../auth"
	"../utils"
)

var (
	URLs = map[string]string{
		"Add": "https://getpocket.com/v3/add",
	}
)

type Response interface{}

type Runable interface {
	Run() (Response, error)
}

type Add struct {
	Runable
	user     *auth.User
	url      string
	title    string
	tags     string
	tweet_id string
	request  utils.HttpRequest
}

type AddResponse struct {
}

func Factory(command string, user *auth.User, targetURL string) Runable {
	switch command {
	case "add":
		request := utils.NewRequest()
		return NewAdd(user, request, targetURL)

	}
	return nil
}

func NewAdd(user *auth.User, request utils.HttpRequest, targetURL string) *Add {
	return &Add{
		user:    user,
		url:     targetURL,
		request: request,
	}
}

func (c *Add) SetUrl(targetURL string) *Add {
	c.url = targetURL
	return c
}

func (c *Add) SetTitle(title string) *Add {
	c.title = title
	return c
}

func (c *Add) SetTags(tags string) *Add {
	c.tags = tags
	return c
}

func (c *Add) SetTweetID(tweet_id string) *Add {
	c.tweet_id = tweet_id
	return c
}

func (c *Add) Run() (Response, error) {
	u := url.Values{}
	u.Add("url", c.url)
	u.Add("consumer_key", "22013-728a24a8a93b5c6de9ca7ba1")
	u.Add("access_token", c.user.AccessToken)
	if "" != c.title {
		u.Add("title", c.title)
	}

	if "" != c.tags {
		u.Add("tags", c.tags)
	}

	if "" != c.tweet_id {
		u.Add("tweet_id", c.tweet_id)
	}

	log.Println(URLs["Add"], u)
	body, err := c.request.Post(URLs["Add"], u)
	if nil != err {
		log.Println("YALLA", err)
		return nil, err
	}
	log.Println("AAAAA", string(body))
	return nil, nil
}
