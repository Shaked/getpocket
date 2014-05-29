package modify

import "encoding/json"

type TagsReplace struct {
	*Protocol
	*ProtocolTags
}

func FactoryTagsReplace(id int, tags []string) *TagsReplace {
	p := &Protocol{Id: id}
	pt := &ProtocolTags{}
	pt.SetTags(tags)
	return NewTagsReplace(p, pt)
}

func NewTagsReplace(p *Protocol, pt *ProtocolTags) *TagsReplace {

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
