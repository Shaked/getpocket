package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Shaked/getpocket/utils"
)

var request = utils.NewRequest()

func TestNew(t *testing.T) {
	_, _ = Factory("consumerKey", "http://www.c.eom")
	_, e := New("consumerKey", ":ww.s.com", request)
	if http.StatusInternalServerError != e.ErrorCode() {
		t.Error(e)
	}

	_, e = New("consumerKey", "http://someurl.org", request)
	if http.StatusInternalServerError != e.ErrorCode() && "HTTPS is required. http://someurl.org" != e.Error() {
		t.Error(e)
	}

	_, e = New("consumerKey", "https://someurl.org", request)
	if nil != e {
		t.Error(e)
	}
}

func TestConnectSuccess(t *testing.T) {
	pocketServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			d := map[string]string{"Code": "123123123"}
			res, _ := json.Marshal(d)
			fmt.Fprint(w, string(res))
		}))

	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	defer pocketServer.Close()
	defer ts.Close()
	mainURL = pocketServer.URL

	a, e := New("consumerKey", ts.URL, request)
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
}
func TestConnectErrors(t *testing.T) {
	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	requests := []struct {
		mainURL  string
		result   string
		errorMsg string
		headers  map[string]string
	}{
		{
			mainURL:  "%s",
			result:   "",
			errorMsg: "invalid json",
			headers:  map[string]string{},
		},
		{
			mainURL:  ":h",
			result:   "result",
			errorMsg: "invalid scheme",
			headers:  map[string]string{},
		},
		{
			mainURL:  "http://some.%s",
			result:   "result",
			errorMsg: "invalid response",
			headers:  map[string]string{},
		},
		{
			mainURL:  "%s",
			result:   "",
			errorMsg: "invalid header errors",
			headers: map[string]string{
				"StatusCode":   strconv.Itoa(http.StatusInternalServerError),
				"X-Error-Code": "132",
				"X-Error":      "Text Error",
			},
		},
	}

	type testResult struct {
		success bool
		msg     string
	}

	ch := make(chan *testResult, len(requests))
	defer close(ch)
	go (func(chan *testResult) {
		for _, r := range requests {
			tr := &testResult{true, ""}
			a, _ := New("consumerKey", ts.URL, request)
			handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				var sc int
				s, ok := r.headers["StatusCode"]
				if ok {
					sc, _ = strconv.Atoi(s)
					delete(r.headers, "StatusCode")
					http.Error(w, "error", sc)
				}

				for hk, hv := range r.headers {
					w.Header().Set(hk, hv)
				}

				if "" != r.result {
					fmt.Fprint(w, r.result)
				}
			})
			pocketServer := httptest.NewServer(handler)
			defer pocketServer.Close()
			mainURL = fmt.Sprintf(r.mainURL, pocketServer.URL)
			c, e := a.Connect()
			if nil == e {
				tr.success = false
				tr.msg = fmt.Sprintf("%s, %s", r.errorMsg, c)
			}
			ch <- tr
		}
	})(ch)

	for i := 0; i < len(requests); i++ {
		tr := <-ch
		if !tr.success {
			t.Error(tr.msg)
		}
	}
}

func TestUserSuccess(t *testing.T) {
	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	pocketServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			d := map[string]string{"username": "shaked", "access_token": "accessToken"}
			res, _ := json.Marshal(d)
			fmt.Fprint(w, string(res))
		}))

	a, _ := New("consumerKey", ts.URL, request)
	mainURL = pocketServer.URL
	user, err := a.User("requestToken")
	if nil != err {
		t.Errorf("invaild user, error: %s", err)
	}

	if "accessToken" != user.AccessToken() || "shaked" != user.Username() {
		t.Errorf("invaild user, user: %#v", user)
	}
}

func TestUserErrors(t *testing.T) {
	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()

	var toClose bool
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if toClose {
			w.Write([]byte("|"))
			v, _ := w.(http.Hijacker)
			conn, _, _ := v.Hijack()
			conn.Close()
		}
	})

	pocketServer := httptest.NewServer(handler)
	defer pocketServer.Close()

	a, _ := New("consumerKey", ts.URL, request)
	mainURL = pocketServer.URL
	r, err := a.User("requestToken")
	if nil == err {
		t.Errorf("json should be invalid, error %s", r)
	}

	toClose = true
	r, err = a.User("requestToken")
	if nil == err {
		t.Errorf("request should be closed, error %s", r)
	}

}

func TestRequestPermissions(t *testing.T) {
	pocketServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a, _ := New("consumerKey", r.URL.Query().Get("redirectURI"), request)
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
