package commands

import (
	"strings"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

var (
	URLs = map[string]string{
		"Add": "https://getpocket.com/v3/add",
	}
)

type Item struct {
	Id             string                `json:"item_id"`
	NormalURL      string                `json:"normal_url"`
	ResolvedId     string                `json:"resolved_id"`
	ResolvedURL    string                `json:"resolved_url"`
	DomainId       string                `json:"domain_id"`
	OriginDomainId string                `json:"origin_domain_id"`
	ResponseCode   string                `json:"response_code"`
	MimeType       string                `json:"mime_type"`
	ContentLength  string                `json:"content_length"`
	Encoding       string                `json:"encoding"`
	DateResolved   string                `json:"date_resolved"`
	DatePublished  string                `json:"date_published"`
	Title          string                `json:"title"`
	Excerpt        string                `json:"excerpt"`
	WordCount      string                `json:"word_count"`
	HasImage       string                `json:"has_image"`
	HasVideo       string                `json:"has_video"`
	IsIndex        string                `json:"is_index"`
	Authors        map[string]ItemAuthor `json:"authors"`
	Images         map[string]ItemImage  `json:"images"`
	Videos         map[string]ItemVideo  `json:"videos"`
}

type ItemImage struct {
	Id      string `json:"item_id"`
	ImageId string `json:"image_id"`
	Src     string `json:"src"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Credit  string `json:"credit"`
}

type ItemVideo struct {
	Id      string `json:"item_id"`
	VideoId string `json:"video_id"`
	Src     string `json:"src"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Type    string `json:"credit"`
	Vid     string `json:"vid"`
	Length  string `json:"length"`
}

type ItemAuthor struct {
	Id   string `json:"author_id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

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

//get pocket returns an empty array instead of an empty object.
func FixJSONArrayToObject(body []byte) []byte {
	newStr := string(body)
	newStr = strings.Replace(newStr, `"videos":[]`, `"videos":{}`, -1)
	newStr = strings.Replace(newStr, `"images":[]`, `"images":{}`, -1)
	newStr = strings.Replace(newStr, `"authors":[]`, `"authors":{}`, -1)
	return []byte(newStr)
}
