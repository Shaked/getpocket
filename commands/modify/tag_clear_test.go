package modify

import "testing"

func TestTagsClearMarshalJSON(t *testing.T) {
	a := FactoryTagsClear(1)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
	}{
		Action: "tags_clear",
		Id:     1,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
	}{
		Action: "tags_clear",
		Id:     1,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
