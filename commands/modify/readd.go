package modify

import "encoding/json"

type Readd struct {
	*Protocol
}

func FactoryReadd(id int) *Readd {
	p := &Protocol{Id: id}
	return NewReadd(p)
}

func NewReadd(p *Protocol) *Readd {
	return &Readd{Protocol: p}
}

func (a *Readd) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "readd",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
