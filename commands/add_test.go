package commands

import (
	"errors"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

func TestNewAdd(t *testing.T) {
	a := NewAdd("target_url")
	if "target_url" != a.URL {
		t.Fail()
	}

	a.SetTitle("title").SetTags("tag1,tags2").SetTweetID("1234")
	user := &auth.User{AccessToken: "access_token", Username: "username"}
	r := &request{ret: "{}"}
	_, err := a.exec(user, "consumerKey", r)
	if nil != err {
		t.Errorf("error %s", err)
	}

	r = &request{err: utils.NewRequestError(1, errors.New("just an error"))}
	_, err = a.exec(user, "consumerKey", r)
	if nil == err {
		t.Fail()
	}

	r = &request{ret: "\n"}
	_, err = a.exec(user, "consumerKey", r)
	if nil == err {
		t.Fail()
	}
}
