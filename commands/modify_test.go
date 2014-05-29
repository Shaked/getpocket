package commands

import (
	"errors"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/commands/modify"
	"github.com/Shaked/getpocket/utils"
)

type ActionStub struct{}

func (a *ActionStub) MarshalJSON() ([]byte, error) {
	return nil, errors.New("error")
}

func TestModifyExec(t *testing.T) {
	user := &auth.User{AccessToken: "access_token", Username: "username"}
	r := &request{ret: "{}"}
	actions := []modify.Action{&ActionStub{}}
	a := NewModify(actions)
	d, err := a.exec(user, "consumerKey", r)
	if nil == err {
		t.Errorf("error %s", d)
	}
	action1 := modify.FactoryFavorite(1)
	action2 := modify.FactoryFavorite(2)
	actions = []modify.Action{action1, action2}
	a = NewModify(actions)
	_, err = a.exec(user, "consumerKey", r)
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
