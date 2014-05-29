package modify

import "testing"

func TestUnfavoriteMarshalJSON(t *testing.T) {
	a := FactoryUnfavorite(1)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
	}{
		Action: "unfavorite",
		Id:     1,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
	}{
		Action: "unfavorite",
		Id:     1,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
