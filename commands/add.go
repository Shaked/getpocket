package commands

import (
	"encoding/json"
	"net/url"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

//@see http://getpocket.com/developer/docs/v3/add
type Add struct {
	command

	URL      string
	title    string
	tags     string
	tweet_id string
}

type AddResponse struct {
	Item   Item
	Status int
}

func NewAdd(consumerKey string, request utils.HttpRequest, targetURL string) *Add {
	a := &Add{
		URL: targetURL,
	}
	a.SetConsumerKey(consumerKey)
	a.SetRequest(request)

	return a
}

// This can be included for cases where an item does not have a title, which is typical for image or PDF URLs.
// If Pocket detects a title from the content of the page, this parameter will be ignored.
func (a *Add) SetTitle(title string) *Add {
	a.title = title
	return a
}

func (a *Add) SetTags(tags string) *Add {
	a.tags = tags
	return a
}

func (a *Add) SetTweetID(tweet_id string) *Add {
	a.tweet_id = tweet_id
	return a
}

func (a *Add) Exec(user auth.Authenticated) (*AddResponse, error) {
	u := url.Values{}

	u.Add("url", a.URL)
	u.Add("consumer_key", a.consumerKey)
	u.Add("access_token", user.AccessToken())

	if "" != a.title {
		u.Add("title", a.title)
	}

	if "" != a.tags {
		u.Add("tags", a.tags)
	}

	if "" != a.tweet_id {
		u.Add("tweet_id", a.tweet_id)
	}

	body, err := a.request.Post(URLs["Add"], u)
	if nil != err {
		return nil, err
	}

	resp := &AddResponse{}
	e := json.Unmarshal(fixJSONArrayToObject(body), resp)
	if nil != e {
		return nil, e
	}
	return resp, nil
}
