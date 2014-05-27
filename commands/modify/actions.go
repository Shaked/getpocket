package modify

import "strings"

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

type Actionable interface{}
type BaseActionTags struct {
	Tags string `json:"tags"`
}
type BaseAction struct {
	Action string `json:"action"`
	Id     int    `json:"item_id"`
	Time   string `json:"timestamp,omitempty"`
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
	switch action {
	case ACTION_ADD:
		a := &Add{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_READD:
		a := &Readd{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_FAVORITE:
		a := &Favorite{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_UNFAVORITE:
		a := &Unfavorite{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_DELETE:
		a := &Delete{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_TAGS_ADD:
		a := &TagsAdd{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_TAGS_REMOVE:
		a := &TagsRemove{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_TAGS_REPLACE:
		a := &TagsReplace{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_TAGS_CLEAR:
		a := &TagsClear{}
		a.Id = id
		a.Action = action
		return a
	case ACTION_TAGS_RENAME:
		a := &TagsRename{}
		a.Id = id
		a.Action = action
		return a
	}
	return nil
}

func (a BaseActionTags) SetTags(tags []string) BaseActionTags {
	a.Tags = strings.Join(tags, ",")
	return a
}
