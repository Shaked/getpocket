package commands

import (
	"errors"
	"testing"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/utils"
)

func TestNewRetrieve(t *testing.T) {
	r := NewRetrieve()

	r.SetContentType(CONTENT_TYPE_VIDEO)
	if CONTENT_TYPE_VIDEO != r.contentType {
		t.Errorf("Content type setter is broken, set to: %s", r.contentType)
	}
	r.SetDetailType(DETAIL_TYPE_SIMPLE)
	if DETAIL_TYPE_SIMPLE != r.detailType {
		t.Errorf("Detail type setter is broken, set to: %s", r.detailType)
	}
	r.SetSort(SORT_SITE)
	if SORT_SITE != r.sort {
		t.Errorf("Sort setter is broken, set to: %s", r.sort)
	}
	r.SetState(STATE_UNREAD)
	if STATE_UNREAD != r.state {
		t.Errorf("State setter is broken, set to: %s", r.state)
	}
	r.SetTag("tag").SetFavorite(true)
	if "tag" != r.tag {
		t.Errorf("Tag setter is broken, set to: %s", r.tag)
	}
	if true != r.favorite {
		t.Errorf("Favorite setter is broken, set to: %s", r.favorite)
	}

	user := &auth.User{AccessToken: "access_token", Username: "username"}
	req := &request{ret: "{}"}
	_, err := r.Exec(user, "consumerKey", req)
	if nil != err {
		t.Errorf("error %s", err)
	}

	req = &request{err: utils.NewRequestError(1, errors.New("just an error"))}
	_, err = r.Exec(user, "consumerKey", req)
	if nil == err {
		t.Fail()
	}

	req = &request{ret: "\n"}
	_, err = r.Exec(user, "consumerKey", req)
	if nil == err {
		t.Fail()
	}
}

func TestSettersError(t *testing.T) {
	r := NewRetrieve()
	var e error
	e = r.SetContentType("contentType")
	if nil == e {
		t.Error("Invalid content type should return an error")
	}
	e = r.SetDetailType("detailType")
	if nil == e {
		t.Error("Invalid detail type should return an error")
	}
	e = r.SetSort("sort")
	if nil == e {
		t.Error("Invalid sort should return an error")
	}
	e = r.SetState("state")
	if nil == e {
		t.Error("Invalid state should return an error")
	}

	r.SetUntagged()
	if TAG_UNTAGGED != r.tag {
		t.Errorf("Tag should be %s when set to untagged", TAG_UNTAGGED)
	}
}
