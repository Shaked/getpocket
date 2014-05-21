package commands

import (
	"encoding/json"
	"errors"
	"net/url"

	"../auth"
	"../utils"
)

var (
	URLs = map[string]string{
		"Add": "https://getpocket.com/v3/add",
	}
)

type Response interface {
}

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

type Add struct {
	Executable
	URL      string
	title    string
	tags     string
	tweet_id string
}

type AddResponse struct {
	Item   Item
	Status int
}

type Item struct {
	Id        string `json:"item_id"`
	NormalURL string `json:"normal_url"`
}

//@see http://getpocket.com/developer/docs/v3/add
func NewAdd(targetURL string) *Add {
	return &Add{
		URL: targetURL,
	}
}

// This can be included for cases where an item does not have a title, which is typical for image or PDF URLs.
// If Pocket detects a title from the content of the page, this parameter will be ignored.
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

func (c *Add) Exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
	if "" == c.URL {
		return nil, errors.New("URL is not defined.")
	}
	u := url.Values{}

	u.Add("url", c.URL)
	u.Add("consumer_key", consumerKey)
	u.Add("access_token", user.AccessToken)

	if "" != c.title {
		u.Add("title", c.title)
	}

	if "" != c.tags {
		u.Add("tags", c.tags)
	}

	if "" != c.tweet_id {
		u.Add("tweet_id", c.tweet_id)
	}

	body, err := request.Post(URLs["Add"], u)
	if nil != err {
		return nil, err
	}

	resp := &AddResponse{}
	e := json.Unmarshal(body, resp)
	if nil != e {
		return nil, e
	}

	return resp, nil
}
