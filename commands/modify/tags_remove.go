package modify

import "encoding/json"

type TagsRemove struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsRemove(id int, tags []string) *TagsRemove {
	p := &Protocol{Id: id}
	pt := &ProtocolTags{}
	pt.SetTags(tags)
	return NewTagsRemove(p, pt)
}

func NewTagsRemove(p *Protocol, pt *ProtocolTags) *TagsRemove {
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
