package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Shaked/getpocket/utils"
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
	RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request)
	Connect() (string, *utils.RequestError)
	User() (*User, *utils.RequestError)
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
	request     utils.HttpRequest
}

type authResponseCode struct {
	Code string `json:"code"`
}

func Factory(consumerKey, redirectURI string) (*Auth, *utils.RequestError) {
	request := utils.NewRequest()
	auth, err := New(consumerKey, redirectURI, request)
	return auth, err
}

//Creates new getpocket auth instance
//consumerKey is available at http://getpocket.com/developers, redirectURI must include full URL
func New(consumerKey, redirectURI string, request utils.HttpRequest) (*Auth, *utils.RequestError) {
	p, err := url.Parse(redirectURI)
	if nil != err {
		return nil, utils.NewRequestError(http.StatusInternalServerError, err)
	}

	if "https" != p.Scheme {
		return nil, utils.NewRequestError(http.StatusInternalServerError, errors.New(
			fmt.Sprintf("Invalid redirectURI, HTTPS is required. %s", p.RawQuery)))
	}

	return &Auth{
		consumerKey: consumerKey,
		redirectURI: redirectURI,
		request:     request,
	}, nil
}

//Connect to GP API using the consumerKey
func (a *Auth) Connect() (string, *utils.RequestError) {
	code, err := a.getCode()
	if nil != err {
		return "", utils.NewRequestError(http.StatusInternalServerError, err)
	}
	return code, nil
}

//Request GP API for permissions
func (a *Auth) RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request) {
	//not checking for error as it is being checked when calling New()
	u, _ := url.Parse(a.redirectURI)
	q := u.Query()
	q.Add("requestToken", requestToken)
	u.RawQuery = q.Encode()
	redirectURL := fmt.Sprintf(
		URLs["RequestTokenUrl"],
		mainURL,
		requestToken,
		url.QueryEscape(u.String()),
	)

	if a.mobile {
		redirectURL += "&mobile=1"
	}

	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

//Make final authentication and retrieve user details (username, access_token)
func (a *Auth) User(requestToken string) (*User, *utils.RequestError) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("code", requestToken)
	body, err := a.request.Post(fmt.Sprintf(URLs["RequestAuthUrl"], mainURL), values)

	if nil != err {
		return nil, err
	}

	user := &User{}
	e := json.Unmarshal(body, user)
	if nil != e {
		return nil, utils.NewRequestError(http.StatusInternalServerError, e)
	}
	return user, nil
}

//@see http://getpocket.com/developer/docs/authentication
func (a *Auth) SetForceMobile(forceMobile bool) {
	a.mobile = forceMobile
}

// private methods
func (a *Auth) getCode() (string, *utils.RequestError) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("redirect_uri", a.redirectURI)
	requestURL := fmt.Sprintf(URLs["RequestUrl"], mainURL)
	body, err := a.request.Post(requestURL, values)
	if nil != err {
		return "", utils.NewRequestError(http.StatusInternalServerError, err)
	}

	res := &authResponseCode{}

	e := json.Unmarshal(body, res)
	if nil != e {
		return "", utils.NewRequestError(http.StatusInternalServerError, e)
	}
	return res.Code, nil
}
