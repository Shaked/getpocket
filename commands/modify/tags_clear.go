package modify

import "encoding/json"

type TagsClear struct {
	*Protocol
}

func FactoryTagsClear(id int) *TagsClear {
	p := &Protocol{Id: id}
	return NewTagsClear(p)
}

func NewTagsClear(p *Protocol) *TagsClear {
	return &TagsClear{Protocol: p}
}

func (a *TagsClear) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "tags_clear",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
