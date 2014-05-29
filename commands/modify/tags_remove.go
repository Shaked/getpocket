package modify

import "encoding/json"

type TagsRemove struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsRemove(id int) *TagsRemove {
	p := &Protocol{Id: id}
	return NewTagsRemove(p)
}

func NewTagsRemove(p *Protocol) *TagsRemove {
	pt := &ProtocolTags{}
	return &TagsRemove{Protocol: p, ProtocolTags: pt}
}

func (a *TagsRemove) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
		*ProtocolTags
	}{
		Action:       "tags_remove",
		Protocol:     a.Protocol,
		ProtocolTags: a.ProtocolTags,
	}

	return json.Marshal(m)
}
