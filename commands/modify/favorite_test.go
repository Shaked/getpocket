package modify

import "testing"

func TestFavoriteMarshalJSON(t *testing.T) {
	a := FactoryFavorite(1)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
	}{
		Action: "favorite",
		Id:     1,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
	}{
		Action: "favorite",
		Id:     1,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
