package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	mainURL = "https://getpocket.com"
	URLs    = map[string]string{
		"RequestUrl":      "%s/v3/oauth/request",
		"RequestTokenUrl": "%s/auth/authorize?request_token=%s&redirect_uri=%s",
		"RequestAuthUrl":  "%s/v3/oauth/authorize",
	}
)

//Authentication methods to connect and use GetPocket API
type Authenticator interface {
	RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request) error
	Connect() (string, error)
	User() (*AuthUser, error)
}

//Optional device settings when using GetPocket API
type DeviceControllable interface {
	SetForceMobile(forceMobile bool)
}

type Auth struct {
	Authenticator
	DeviceControllable
	consumerKey string
	redirectURI string
	mobile      bool
}

type authResponseCode struct {
	Code string `json:"code"`
}

//Creates new gogetpocket auth instance
//consumerKey is available at http://getpocket.com/developers, redirectURI must include full URL
func New(consumerKey, redirectURI string) *Auth {
	return &Auth{
		consumerKey: consumerKey,
		redirectURI: redirectURI,
	}
}

//Connect to GP API using the consumerKey
func (a *Auth) Connect() (string, error) {
	code, err := a.getCode()
	return code, err
}

//Request GP API for permissions
func (a *Auth) RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request) error {
	u, err := url.Parse(a.redirectUri)
	if nil != err {
		return err
	}
	q := u.Query()
	q.Add("requestToken", requestToken)
	u.RawQuery = q.Encode()
	redirectUrl := fmt.Sprintf(
		URLs["RequestTokenUrl"],
		mainURL,
		requestToken,
		url.QueryEscape(u.String()),
	)

	if a.mobile {
		urls["RequestTokenUrl"] += "&mobile=1"
	}

	http.Redirect(w, r, redirectUrl, 302)
	return nil
}

//Make final authentication and retrieve user details (username, access_token)
func (a *Auth) User(requestToken string) (*AuthUser, error) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("code", requestToken)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Accept":     "application/json",
	}

	body, err := a.post(URLs["RequestAuthUrl"], mainURL, values, headers)

	if nil != err {
		return nil, err
	}

	user := &AuthUser{}
	err = json.Unmarshal(body, user)
	return user, err
}

func (a *Auth) SetForceMobile(forceMobile bool) {
	a.mobile = forceMobile
}

// private methods

func (a *Auth) getCode() (string, error) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("redirect_uri", a.redirectURI)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Accept":     "application/json",
	}

	body, err := a.post(URLs["RequestUrl"], mainURL, values, headers)
	if nil != err {
		return "", err
	}

	res := &authResponseCode{}
	err = json.Unmarshal(body, res)
	return res.Code, err
}

func (a *Auth) post(url string, values url.Values, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if nil != err {
		return nil, err
	}
	for header, value := range headers {
		r.Header.Set(header, value)
	}

	resp, err := client.Do(r)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		xErrorCode, _ := strconv.ParseInt(resp.Header.Get("X-Error-Code"), 10, 0)
		xErrorText := resp.Header.Get("X-Error")
		return nil, NewAuthError(int(xErrorCode), xErrorText)
	}
	return body, err
}
