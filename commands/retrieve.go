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
	favorite    bool
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
	List     map[string]RetrieveItem `json:"list"`
	Status   int                     `json:"status"`
	Complete int                     `json:"complete"`
}

type RetrieveItem struct {
	Id            string                `json:"item_id"`
	ResolvedId    string                `json:"resolved_id"`
	GivenURL      string                `json:"given_url"`
	ResolvedURL   string                `json:"resolved_url"`
	GivenTitle    string                `json:"given_title"`
	ResolvedTitle string                `json:"resolved_title"`
	Favorite      string                `json:"favorite"`
	Status        string                `json:"status"`
	Excerpt       string                `json:"excerpt"`
	IsArticle     string                `json:"is_article"`
	HasImage      string                `json:"has_image"`
	HasVideo      string                `json:"has_video"`
	WordCount     string                `json:"word_count"`
	TimeAdded     string                `json:"time_added"`
	TimeUpdated   string                `json:"time_updated"`
	TimeRead      string                `json:"time_read"`
	TimeFavorited string                `json:"time_favorited"`
	SortId        int                   `json:"sort_id"`
	IsIndex       string                `json:"is_index"`
	Tags          map[string]ItemTag    `json:"tags"`
	Authors       map[string]ItemAuthor `json:"authors"`
	Images        map[string]ItemImage  `json:"images"`
	Videos        map[string]ItemVideo  `json:"videos"`
}

func NewRetrieve() *Retrieve {
	return &Retrieve{}
}

func (c *Retrieve) SetState(state string) error {
	if !states[state] {
		return errors.New(fmt.Sprintf("State %s does not exist", state))
	}
	c.state = state
	return nil
}

func (c *Retrieve) SetFavorite(favorite bool) *Retrieve {
	c.favorite = favorite
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

func (c *Retrieve) SetContentType(contentType string) error {
	if !contentTypes[contentType] {
		return errors.New(fmt.Sprintf("ContentType %s does not exist", contentType))
	}
	c.contentType = contentType
	return nil
}

func (c *Retrieve) SetSort(sort string) error {
	if !sorts[sort] {
		return errors.New(fmt.Sprintf("sort %s does not exist", sort))
	}
	c.sort = sort
	return nil
}

func (c *Retrieve) SetDetailType(detailType string) error {
	if !detailTypes[detailType] {
		return errors.New(fmt.Sprintf("detailType %s does not exist", detailType))
	}
	c.detailType = detailType
	return nil
}

func (c *Retrieve) exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
	u := url.Values{}
	u.Add("consumer_key", consumerKey)
	u.Add("access_token", user.AccessToken)

	if "" != c.state {
		u.Add("state", c.state)
	}

	if c.favorite {
		u.Add("favorite", "1")
	}

	if "" != c.tag {
		u.Add("tag", c.tag)
	}

	if "" != c.contentType {
		u.Add("contentType", c.contentType)
	}

	if "" != c.sort {
		u.Add("sort", c.sort)
	}

	if "" != c.detailType {
		u.Add("detailType", c.detailType)
	}

	body, err := request.Post(URLs["Retrieve"], u)
	if nil != err {
		return nil, err
	}

	resp := &RetrieveResponse{}
	e := json.Unmarshal(fixJSONArrayToObject(body), resp)
	if nil != e {
		return nil, e
	}
	return resp, nil
}
