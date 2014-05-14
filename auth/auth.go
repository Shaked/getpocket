package auth

import (
	"encoding/json"
	"errors"
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
	RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request) *AuthError
	Connect() (string, *AuthError)
	User() (*AuthUser, *AuthError)
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
func (a *Auth) Connect() (string, *AuthError) {
	code, err := a.getCode()
	if nil != err {
		return "", NewAuthError(http.StatusInternalServerError, err)
	}
	return code, nil
}

//Request GP API for permissions
func (a *Auth) RequestPermissions(requestToken string, w http.ResponseWriter, r *http.Request) *AuthError {
	u, err := url.Parse(a.redirectURI)
	if nil != err {
		return NewAuthError(http.StatusInternalServerError, err)
	}
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

	http.Redirect(w, r, redirectURL, 302)
	return nil
}

//Make final authentication and retrieve user details (username, access_token)
func (a *Auth) User(requestToken string) (*AuthUser, *AuthError) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("code", requestToken)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Accept":     "application/json",
	}

	body, err := a.post(fmt.Sprintf(URLs["RequestAuthUrl"], mainURL), values, headers)

	if nil != err {
		return nil, err
	}

	user := &AuthUser{}
	e := json.Unmarshal(body, user)
	if nil != err {
		return nil, NewAuthError(http.StatusInternalServerError, e)
	}
	return user, nil
}

//@see http://getpocket.com/developer/docs/authentication
func (a *Auth) SetForceMobile(forceMobile bool) {
	a.mobile = forceMobile
}

// private methods
func (a *Auth) getCode() (string, *AuthError) {
	values := make(url.Values)
	values.Set("consumer_key", a.consumerKey)
	values.Set("redirect_uri", a.redirectURI)
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Accept":     "application/json",
	}

	body, err := a.post(fmt.Sprintf(URLs["RequestUrl"], mainURL), values, headers)
	if nil != err {
		return "", NewAuthError(http.StatusInternalServerError, err)
	}

	res := &authResponseCode{}
	e := json.Unmarshal(body, res)
	if nil != err {
		return "", NewAuthError(http.StatusInternalServerError, e)
	}
	return res.Code, nil
}

func (a *Auth) post(url string, values url.Values, headers map[string]string) ([]byte, *AuthError) {
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if nil != err {
		return nil, NewAuthError(http.StatusInternalServerError, err)
	}
	for header, value := range headers {
		r.Header.Set(header, value)
	}

	resp, err := client.Do(r)
	if nil != err {
		return nil, NewAuthError(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, NewAuthError(http.StatusInternalServerError, err)
	}
	if resp.StatusCode != http.StatusOK {
		xErrorCode, _ := strconv.ParseInt(resp.Header.Get("X-Error-Code"), 10, 0)
		xErrorText := resp.Header.Get("X-Error")
		return nil, NewAuthError(int(xErrorCode), errors.New(xErrorText))
	}
	return body, nil
}
