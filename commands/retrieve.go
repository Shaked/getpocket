package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

const (
	STATE_UNREAD  = "unread"
	STATE_ARCHIVE = "archive"
	STATE_ALL     = "all"

	TAG_UNTAGGED = "_untagged_"

	CONTENT_TYPE_ARTICLE = "article"
	CONTENT_TYPE_VIDEO   = "video"
	CONTENT_TYPE_IMAGE   = "image"

	SORT_NEWEST = "newest"
	SORT_OLDEST = "oldest"
	SORT_TITLE  = "title"
	SORT_SITE   = "site"

	DETAIL_TYPE_SIMPLE   = "simple"
	DETAIL_TYPE_COMPLETE = "complete"
)

var (
	states = map[string]bool{
		STATE_ALL:     true,
		STATE_ARCHIVE: true,
		STATE_UNREAD:  true,
	}

	contentTypes = map[string]bool{
		CONTENT_TYPE_ARTICLE: true,
		CONTENT_TYPE_IMAGE:   true,
		CONTENT_TYPE_VIDEO:   true,
	}

	sorts = map[string]bool{
		SORT_NEWEST: true,
		SORT_OLDEST: true,
		SORT_SITE:   true,
		SORT_TITLE:  true,
	}

	detailTypes = map[string]bool{
		DETAIL_TYPE_COMPLETE: true,
		DETAIL_TYPE_SIMPLE:   true,
	}
)

//@see http://getpocket.com/developer/docs/v3/retrieve
type Retrieve struct {
	Executable
	state       string
	favorite    int
	tag         string
	contentType string
	sort        string
	detailType  string
	search      string
	domain      string
	since       time.Time
	count       int
	offset      int
}

type RetrieveResponse struct {
	Item   Item
	Status int
}

func NewRetrieve() *Retrieve {
	return &Retrieve{}
}

func (c *Retrieve) SetState(state string) (*Retrieve, error) {
	if !states[state] {
		return nil, errors.New(fmt.Sprintf("State %s does not exist", state))
	}
	c.state = state
	return c, nil
}

func (c *Retrieve) SetFavorite(favorite bool) *Retrieve {
	c.favorite = int(favorite)
	return c
}

func (c *Retrieve) SetTag(tag string) *Retrieve {
	c.tag = tag
	return c
}

func (c *Retrieve) SetUntagged() *Retrieve {
	c.SetTag(TAG_UNTAGGED)
	return c
}

func (c *Retrieve) SetContentType(contentType string) (*Retrieve, error) {
	if !contentTypes[contentType] {
		return nil, errors.New(fmt.Sprintf("ContentType %s does not exist", contentType))
	}
	c.contentType = contentType
	return c, nil
}

func (c *Retrieve) Exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
	u := url.Values{}

	u.Retrieve("url", c.URL)
	u.Retrieve("consumer_key", consumerKey)
	u.Retrieve("access_token", user.AccessToken)

	if "" != c.title {
		u.Retrieve("title", c.title)
	}

	if "" != c.tags {
		u.Retrieve("tags", c.tags)
	}

	if "" != c.tweet_id {
		u.Retrieve("tweet_id", c.tweet_id)
	}

	body, err := request.Post(URLs["Retrieve"], u)
	if nil != err {
		return nil, err
	}

	resp := &RetrieveResponse{}
	e := json.Unmarshal(FixJSONArrayToObject(body), resp)
	if nil != e {
		return nil, e
	}
	return resp, nil
}
