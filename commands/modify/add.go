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

func (a *Add) SetRefId(refId int) *Add {
	a.RefId = refId
	return a
}

func (a *Add) SetTitle(title string) *Add {
	a.Title = title
	return a
}

func (a *Add) SetURL(url string) *Add {
	a.URL = url
	return a
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
		Protocol:     a.Protocol,
		ProtocolTags: a.ProtocolTags,
	}

	if 0 < a.RefId {
		m.RefId = a.RefId
	}

	if "" != a.Title {
		m.Title = a.Title
	}

	if "" != a.URL {
		m.URL = a.URL
	}

	return json.Marshal(m)
}
