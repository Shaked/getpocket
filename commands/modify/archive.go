package modify

import "encoding/json"

type Archive struct {
	*Protocol
}

func FactoryArchive(id int) *Archive {
	p := &Protocol{Id: id}
	return NewArchive(p)
}

func NewArchive(p *Protocol) *Archive {
	return &Archive{Protocol: p}
}

func (a *Archive) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "archive",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
