package modify

import "encoding/json"

type TagsAdd struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsAdd(id int) *TagsAdd {
	p := &Protocol{Id: id}
	return NewTagsAdd(p)
}

func NewTagsAdd(p *Protocol) *TagsAdd {
	pt := &ProtocolTags{}
	return &TagsAdd{Protocol: p, ProtocolTags: pt}
}

func (a *TagsAdd) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
		*ProtocolTags
	}{
		Action:       "tags_add",
		Protocol:     a.Protocol,
		ProtocolTags: a.ProtocolTags,
	}

	return json.Marshal(m)
}
