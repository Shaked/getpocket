package modify

import "encoding/json"

type TagsRename struct {
	*Protocol
	OldTag string
	NewTag string
}

func FactoryTagsRename(id int, oldTag, newTag string) *TagsRename {
	p := &Protocol{Id: id}
	return NewTagsRename(p, oldTag, newTag)
}

func NewTagsRename(p *Protocol, oldTag, newTag string) *TagsRename {
	return &TagsRename{Protocol: p, OldTag: oldTag, NewTag: newTag}
}

func (a *TagsRename) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
		OldTag string `json:"old_tag"`
		NewTag string `json:"new_tag"`
	}{
		Action:   "tags_rename",
		Protocol: a.Protocol,
		OldTag:   a.OldTag,
		NewTag:   a.NewTag,
	}

	return json.Marshal(m)
}
