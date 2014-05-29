package modify

import "testing"

func TestTagsAddMarshalJSON(t *testing.T) {
	a := FactoryTagsAdd(1, []string{`tag1`, `tag2`})
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Tags   string `json:"tags"`
	}{
		Action: "tags_add",
		Id:     1,
		Tags:   `tag1,tag2`,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
		Tags   string `json:"tags"`
	}{
		Action: "tags_add",
		Id:     1,
		Tags:   `tag1,tag2`,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
