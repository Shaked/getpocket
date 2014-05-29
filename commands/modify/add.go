package modify

import "encoding/json"

type Add struct {
	*Protocol
	*ProtocolTags
	RefId int
	Title string
	URL   string
}

func FactoryAdd(id int) *Add {
	p := &Protocol{Id: id}
	return NewAdd(p)
}

func NewAdd(p *Protocol) *Add {
	pt := &ProtocolTags{}
	return &Add{Protocol: p, ProtocolTags: pt}
}

func (a *Add) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
		*ProtocolTags
		RefId int    `json:"ref_id,omitempty"`
		Title string `json:"title,omitempty"`
		URL   string `json:"url,omitempty"`
	}{
		Action:       "add",
		Title:        a.Title,
		RefId:        a.RefId,
		URL:          a.URL,
		Protocol:     a.Protocol,
		ProtocolTags: a.ProtocolTags,
	}
	return json.Marshal(m)
}
