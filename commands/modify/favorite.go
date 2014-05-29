package modify

import "encoding/json"

type Favorite struct {
	*Protocol
}

func FactoryFavorite(id int) *Favorite {
	p := &Protocol{Id: id}
	return NewFavorite(p)
}

func NewFavorite(p *Protocol) *Favorite {
	return &Favorite{Protocol: p}
}

func (a *Favorite) MarshalJSON() ([]byte, error) {
	m := struct {
		Action string `json:"action"`
		*Protocol
	}{
		Action:   "favorite",
		Protocol: a.Protocol,
	}

	return json.Marshal(m)
}
