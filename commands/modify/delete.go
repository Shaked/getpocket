package modify

import "encoding/json"

type Delete struct {
	*Protocol
}

func FactoryDelete(id int) *Delete {
	p := &Protocol{Id: id}
	return NewDelete(p)
}

func NewDelete(p *Protocol) *Delete {
	return &Delete{Protocol: p}
}

func (a *Delete) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "delete",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
