package modify

import "encoding/json"

type Unfavorite struct {
	*Protocol
}

func FactoryUnfavorite(id int) *Unfavorite {
	p := &Protocol{Id: id}
	return NewUnfavorite(p)
}

func NewUnfavorite(p *Protocol) *Unfavorite {
	return &Unfavorite{Protocol: p}
}

func (a *Unfavorite) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "unfavorite",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
