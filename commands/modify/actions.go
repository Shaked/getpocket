package modify

import (
	"encoding/json"
	"strings"
)

type Actionable interface {
	Parse() (string, error)
}

type Add struct {
	Action string `json:"action"`
	Id     int    `json:"item_id"`
	RefId  int    `json:"ref_id,omitempty"`
	Tags   string `json:"tags,omitempty"`
	Time   string `json:"timestamp,omitempty"`
	Title  string `json:"title,omitempty"`
	URL    string `json:"url,omitempty"`
}

// type Archive struct {
// 	Action
// }

// type Readd struct {
// 	Action
// }

// type Favorite struct {
// 	Action
// }

func Factory(action string, id int) Actionable {
	switch action {
	case "add":
		return &Add{Action: "add", Id: id}
	}
	return nil
}

func (a *Add) Parse() (string, error) {
	jsonString, e := json.Marshal(a)
	if nil != e {
		return "", e
	}
	return string(jsonString), nil
}

func (a *Add) SetTags(tags []string) *Add {
	a.Tags = strings.Join(tags, ",")
	return a
}
