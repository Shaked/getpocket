package modify

import "testing"

func TestReaddMarshalJSON(t *testing.T) {
	a := FactoryReadd(1)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
	}{
		Action: "readd",
		Id:     1,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
	}{
		Action: "readd",
		Id:     1,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
