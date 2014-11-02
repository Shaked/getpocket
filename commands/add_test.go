package commands

import (
	"errors"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

func TestNewAdd(t *testing.T) {
	r := &request{ret: "{}"}
	a := NewAdd("consumer-key", r, "target_url")
	if "target_url" != a.URL {
		t.Fail()
	}

	a.SetTitle("title").SetTags("tag1,tags2").SetTweetID("1234")
	user := &auth.User{AccessToken: "access_token", Username: "username"}
	_, err := a.Exec(user)
	if nil != err {
		t.Errorf("error %s", err)
	}

	r = &request{err: utils.NewRequestError(1, errors.New("just an error"))}
	a = NewAdd("consumer-key", r, "target_url")
	_, err = a.Exec(user)
	if nil == err {
		t.Fail()
	}

	r = &request{ret: "\n"}
	a = NewAdd("consumer-key", r, "target_url")
	_, err = a.Exec(user)
	if nil == err {
		t.Fail()
	}
}
