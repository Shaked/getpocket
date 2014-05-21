package commands

import (
	"errors"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

type C struct {
	resp Response
	err  error
}

func (c *C) Exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
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
