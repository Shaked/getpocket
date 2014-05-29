package modify

import "encoding/json"

type TagsAdd struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsAdd(id int, tags []string) *TagsAdd {
	p := &Protocol{Id: id}
	pt := &ProtocolTags{}
	pt.SetTags(tags)
	return NewTagsAdd(p, pt)
}

func NewTagsAdd(p *Protocol, pt *ProtocolTags) *TagsAdd {
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
