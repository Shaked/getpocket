package commands

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/Shaked/getpocket/auth"
	"github.com/Shaked/getpocket/commands/modify"
	"github.com/Shaked/getpocket/utils"
)

//@see http://getpocket.com/developer/docs/v3/add
type Modify struct {
	Executable
	actions []modify.Actionable
}

type ModifyResponse struct {
	ActionResults []bool `json:"action_results"`
	Status        int    `json:"status"`
}

func NewModify(actions []modify.Actionable) *Modify {
	return &Modify{
		actions: actions,
	}
}

func (c *Modify) exec(user *auth.User, consumerKey string, request utils.HttpRequest) (Response, error) {
	u := url.Values{}
	u.Add("consumer_key", consumerKey)
	u.Add("access_token", user.AccessToken)

	// actions := make([]string, len(c.actions))
	// for _, action := range c.actions {
	// 	p, e := action.Parse()
	// 	if nil != e {
	// 		log.Println("E!!!!", e)
	// 	}
	// 	actions = append(actions, p)
	// }
	// log.Fatal("ACTIONS", actions)
	actions, e := json.Marshal(c.actions)
	if nil != e {
		log.Printf("errorrr %s", e)
	}
	u.Add("actions", string(actions))

	body, err := request.Post(URLs["Modify"], u)
	log.Println("SOMETIEST", string(body), err)
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
