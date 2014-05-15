package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	_, e := New("consumerKey", "redirectURI")
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
	pocketServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d := map[string]string{"code": "123123123"}
		res, _ := json.Marshal(d)
		fmt.Fprint(w, string(res))
	}))
	defer pocketServer.Close()

	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {

			}))
	defer ts.Close()

	mainURL = pocketServer.URL

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

	mainURL = "does not exist url"
	c, e = a.Connect()
	if nil == e {
		t.Errorf("Should return an error when URL is invalid, code: %s", c)
	}
}
