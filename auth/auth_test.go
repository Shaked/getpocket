package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	_, e := New("consumerKey", ":ww.s.com")
	if http.StatusInternalServerError != e.ErrorCode() {
		t.Error(e)
	}

	_, e = New("consumerKey", "http://someurl.org")
	if http.StatusInternalServerError != e.ErrorCode() && "HTTPS is required. http://someurl.org" != e.Error() {
		t.Error(e)
	}

	_, e = New("consumerKey", "https://someurl.org")
	if nil != e {
		t.Error(e)
	}
}

func TestConnect(t *testing.T) {
	pocketServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if "" != r.URL.Query().Get("json") {
				d := map[string]string{"Code": "123123123"}
				res, _ := json.Marshal(d)
				fmt.Fprint(w, string(res))
			}

		}))

	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	defer pocketServer.Close()
	defer ts.Close()

	mainURL = pocketServer.URL + "?json=1&"

	a, e := New("consumerKey", ts.URL)
	if nil != e {
		t.Errorf("New function failed, error: %s", e)
	}

	c, e := a.Connect()
	if nil != e {
		t.Errorf("Connect method failed, error: %s", e)
	}

	if "123123123" != c {
		t.Errorf("Wrong API code returned: %s", c)
	}

	mainURL = pocketServer.URL + `?json=&`
	c, e = a.Connect()
	if nil == e {
		t.Errorf("Connect method passed, should return an error, actual: %s", c)
	}

	mainURL = "does not exist url"
	c, e = a.Connect()
	if nil == e {
		t.Errorf("Should return an error when URL is invalid, code: %s", c)
	}
}

func TestRequestPermissions(t *testing.T) {
	pocketServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, _ := New("consumerKey", r.URL.Query().Get("redirectURI"))
		a.SetForceMobile(true)
		a.RequestPermissions("SomeRequestToken", w, r)
	}))

	defer pocketServer.Close()
	defer ts.Close()
	mainURL = pocketServer.URL

	//make sure redirect works
	client := newClient()
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		requestToken := req.URL.Query().Get("request_token")
		if "SomeRequestToken" != requestToken {
			t.Errorf("Wrong request_token, actual: %s", requestToken)
		}
		return nil
	}

	//call request permission url
	req, e := client.Get(ts.URL + "?redirectURI=" + ts.URL)
	if nil != e {
		t.Fatalf("Get request failed, error: %s", e)
	}
	defer req.Body.Close()

}

func newClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
