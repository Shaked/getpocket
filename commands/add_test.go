package commands

import (
	"errors"
	"net/url"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

type request struct {
	ret string
	err *utils.RequestError
}

func (r *request) Post(url string, values url.Values) ([]byte, *utils.RequestError) {
	return []byte(r.ret), r.err
}
func TestNewAdd(t *testing.T) {
	a := NewAdd("target_url")
	if "target_url" != a.URL {
		t.Fail()
	}

	a.SetTitle("title").SetTags("tag1,tags2").SetTweetID("1234")
	user := &auth.User{AccessToken: "access_token", Username: "username"}
	r := &request{ret: "{}"}
	_, err := a.Exec(user, "consumerKey", r)
	if nil != err {
		t.Errorf("error %s", err)
	}

	r = &request{err: utils.NewRequestError(1, errors.New("just an error"))}
	_, err = a.Exec(user, "consumerKey", r)
	if nil == err {
		t.Fail()
	}

	r = &request{ret: "\n"}
	_, err = a.Exec(user, "consumerKey", r)
	if nil == err {
		t.Fail()
	}
}