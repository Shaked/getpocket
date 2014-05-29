package modify

import "testing"

func TestTagsRenameMarshalJSON(t *testing.T) {
	a := FactoryTagsRename(1, `tag1`, `tag2`)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		OldTag string `json:"old_tag"`
		NewTag string `json:"new_tag"`
	}{
		Action: "tags_rename",
		Id:     1,
		OldTag: `tag1`,
		NewTag: `tag2`,
	}

	ValidateJSONs(t, a, m1)

	a.SetTS("time")
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
		OldTag string `json:"old_tag"`
		NewTag string `json:"new_tag"`
	}{
		Action: "tags_rename",
		Id:     1,
		OldTag: `tag1`,
		NewTag: `tag2`,
		Time:   "time",
	}

	ValidateJSONs(t, a, m2)
}
