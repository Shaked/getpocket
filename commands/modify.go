package commands

import (
	"encoding/json"
	"net/url"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/commands/modify"
	"github.com/Shaked/getpocket/utils"
)

//@see http://getpocket.com/developer/docs/v3/add
type Modify struct {
	command

	actions []modify.Action
}

type ModifyResponse struct {
	ActionResults []bool `json:"action_results"`
	Status        int    `json:"status"`
}

func NewModify(consumerKey string, request utils.HttpRequest, actions []modify.Action) *Modify {
	m := &Modify{actions: actions}
	m.SetConsumerKey(consumerKey)
	m.SetRequest(request)
	return m
}

func (c *Modify) Exec(user *auth.User) (*ModifyResponse, error) {
	u := url.Values{}
	u.Add("consumer_key", c.consumerKey)
	u.Add("access_token", user.AccessToken)

	actions, e := json.Marshal(c.actions)
	if nil != e {
		return nil, e
	}
	u.Add("actions", string(actions))
	body, err := c.request.Post(URLs["Modify"], u)
	if nil != err {
		return nil, err
	}

	resp := &ModifyResponse{}
	e = json.Unmarshal(fixJSONArrayToObject(body), resp)
	if nil != e {
		return nil, e
	}
	return resp, nil
}
