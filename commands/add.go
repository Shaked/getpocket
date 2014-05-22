package commands

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

//@see http://getpocket.com/developer/docs/v3/add
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
	e := json.Unmarshal(FixJSONArrayToObject(body), resp)
	if nil != e {
		return nil, e
	}

	// switch resp.Item.Videos.(type) {
	// case map[string]interface{}:
	// 	m := resp.Item.Videos.(map[string]interface{})
	// 	for key, value := range m {
	// 		v := value.(map[string]interface{})
	// 		m[key] = ItemVideo{
	// 			Id:      v["item_id"].(string),
	// 			VideoId: v["video_id"].(string),
	// 			Src:     v["src"].(string),
	// 			Width:   v["width"].(string),
	// 			Height:  v["height"].(string),
	// 			Type:    v["type"].(string),
	// 			Vid:     v["vid"].(string),
	// 			Length:  v["length"].(string),
	// 		}
	// 	}
	// 	resp.Item.Videos = m
	// case []interface{}:
	// 	resp.Item.Videos = nil
	// }

	log.Printf("%#v	", resp.Item.Videos)

	return resp, nil
}
