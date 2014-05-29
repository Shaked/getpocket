package modify

import "encoding/json"

type TagsReplace struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsReplace(id int) *TagsReplace {
	p := &Protocol{Id: id}
	return NewTagsReplace(p)
}

func NewTagsReplace(p *Protocol) *TagsReplace {
	pt := &ProtocolTags{}
	return &TagsReplace{Protocol: p, ProtocolTags: pt}
}

func (a *TagsReplace) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
		*ProtocolTags
	}{
		Action:       "tags_replace",
		Protocol:     a.Protocol,
		ProtocolTags: a.ProtocolTags,
	}

	return json.Marshal(m)
}
