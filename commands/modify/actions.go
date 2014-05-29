package modify

import (
	"encoding/json"
	"strings"
)

type Action interface {
	json.Marshaler
}

type ProtocolTags struct {
	Tags string `json:"tags,omitempty"`
}

type Protocol struct {
	Id   int    `json:"item_id"`
	Time string `json:"timestamp,omitempty"`
}

func (a *ProtocolTags) SetTags(tags []string) *ProtocolTags {
	a.Tags = strings.Join(tags, ",")
	return a
}

func (a *Protocol) SetTS(ts string) *Protocol {
	a.Time = ts
	return a
}
