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

type C struct {
	resp Response
	err  error
}

func (c *C) exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
	return c.resp, c.err
}

func TestNew(t *testing.T) {
	user := &auth.User{}
	c := New(user, "consumerKey")

	if "consumerKey" != c.consumerKey {
		t.Fail()
	}

	if user != c.user {
		t.Fail()
	}
}

func TestExec(t *testing.T) {
	user := &auth.User{}
	c := New(user, "consumerKey")
	stub := struct{}{}
	command := &C{resp: stub, err: errors.New("Error")}
	resp, err := c.Exec(command)
	if "Error" != err.Error() {
		t.Errorf("Error: %s", err)
	}

	if stub != resp {
		t.Errorf("Invalid response: %s", resp)
	}
}

func TestfixJSONArrayToObject(t *testing.T) {
	// This
	apiResult := []byte(`{"item_id":"1", "videos":[],"authors":[],"images":[]}`)
	expected := []byte(`{"item_id":"1", "videos":{},"authors":{},"images":{}}`)
	actual := fixJSONArrayToObject(apiResult)
	if string(expected) != string(actual) {
		t.Errorf("Actual value is worng %s", string(actual))
	}

}
