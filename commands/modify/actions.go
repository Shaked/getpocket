package modify

import (
	"errors"
	"fmt"
	"strings"
)

const (
	ACTION_ADD          = "add"
	ACTION_FAVORITE     = "favorite"
	ACTION_UNFAVORITE   = "unfavorite"
	ACTION_READD        = "readd"
	ACTION_DELETE       = "delete"
	ACTION_TAGS_ADD     = "tags_add"
	ACTION_TAGS_REMOVE  = "tags_remove"
	ACTION_TAGS_REPLACE = "tags_replace"
	ACTION_TAGS_CLEAR   = "tags_clear"
	ACTION_TAGS_RENAME  = "tags_rename"
)

var (
	availableActions = map[string]bool{
		ACTION_ADD:          true,
		ACTION_FAVORITE:     true,
		ACTION_UNFAVORITE:   true,
		ACTION_READD:        true,
		ACTION_DELETE:       true,
		ACTION_TAGS_ADD:     true,
		ACTION_TAGS_REMOVE:  true,
		ACTION_TAGS_REPLACE: true,
		ACTION_TAGS_CLEAR:   true,
		ACTION_TAGS_RENAME:  true,
	}
)

type Actionable interface {
	setAction(action string) error
	setId(id int)
}

type BaseActionTags struct {
	Tags string `json:"tags"`
}

type BaseAction struct {
	Action string `json:"action"`
	Id     int    `json:"item_id"`
	Time   string `json:"timestamp,omitempty"`
}

func (a *BaseAction) setAction(action string) error {
	if availableActions[action] {
		return errors.New(fmt.Sprintf("Action %s does not exist.", action))
	}
	a.Action = action
	return nil
}
func (a *BaseAction) setId(id int) {
	a.Id = id
}

type Add struct {
	BaseAction
	BaseActionTags
	RefId int    `json:"ref_id,omitempty"`
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

type Archive struct {
	BaseAction
}

type Readd struct {
	BaseAction
}

type Favorite struct {
	BaseAction
}

type Unfavorite struct {
	BaseAction
}

type Delete struct {
	BaseAction
}

type TagsAdd struct {
	BaseAction
	BaseActionTags
}

type TagsRemove struct {
	BaseAction
	BaseActionTags
}

type TagsReplace struct {
	BaseAction
	BaseActionTags
}

type TagsClear struct {
	BaseAction
}

type TagsRename struct {
	BaseAction
	OldTag string `json:"old_tag"`
	NewTag string `json:"new_tag"`
}

func Factory(action string, id int) Actionable {
	var a Actionable
	switch action {
	case ACTION_ADD:
		a = &Add{}
		break
	case ACTION_READD:
		a = &Readd{}
		break
	case ACTION_FAVORITE:
		a = &Favorite{}
		break
	case ACTION_UNFAVORITE:
		a = &Unfavorite{}
		break
	case ACTION_DELETE:
		a = &Delete{}
		break
	case ACTION_TAGS_ADD:
		a = &TagsAdd{}
		break
	case ACTION_TAGS_REMOVE:
		a = &TagsRemove{}
		break
	case ACTION_TAGS_REPLACE:
		a = &TagsReplace{}
		break
	case ACTION_TAGS_CLEAR:
		a = &TagsClear{}
		break
	case ACTION_TAGS_RENAME:
		a = &TagsRename{}
		break
	}

	a.setId(id)
	a.setAction(action)
	return a
}

func (a *BaseActionTags) SetTags(tags []string) *BaseActionTags {
	a.Tags = strings.Join(tags, ",")
	return a
}
