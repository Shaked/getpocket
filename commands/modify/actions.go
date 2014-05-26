package modify

import "strings"

const (
	ACTION_ADD        = "add"
	ACTION_FAVORITE   = "favorite"
	ACTION_UNFAVORITE = "unfavorite"
)

type Actionable interface{}

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

type Favorite struct {
	Action string `json:"action"`
	Id     int    `json:"item_id"`
	Time   string `json:"timestamp,omitempty"`
}

type Unfavorite struct {
	Action string `json:"action"`
	Id     int    `json:"item_id"`
	Time   string `json:"timestamp,omitempty"`
}

func Factory(action string, id int) Actionable {
	switch action {
	case "add":
		return &Add{Action: "add", Id: id}
	case ACTION_FAVORITE:
		return &Favorite{Action: ACTION_FAVORITE, Id: id}
	case ACTION_UNFAVORITE:
		return &Unfavorite{Action: ACTION_UNFAVORITE, Id: id}
	}
	return nil
}

func (a *Add) SetTags(tags []string) *Add {
	a.Tags = strings.Join(tags, ",")
	return a
}
