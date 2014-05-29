package modify

import "testing"

func TestAddMarshalJSON(t *testing.T) {
	a := FactoryAdd(1)
	m1 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
	}{
		Action: "add",
		Id:     1,
	}
	ValidateJSONs(t, a, m1)

	now := "123123123"
	a.SetRefId(2).SetTitle("title").SetURL("url")
	a.SetTags([]string{"tag1", "tag2"})
	a.SetTS(now)
	m2 := struct {
		Action string `json:"action"`
		Id     int    `json:"item_id"`
		Time   string `json:"timestamp"`
		Tags   string `json:"tags"`
		RefId  int    `json:"ref_id"`
		Title  string `json:"title"`
		URL    string `json:"url"`
	}{
		Action: "add",
		Id:     1,
		Time:   now,
		Tags:   "tag1,tag2",
		RefId:  a.RefId,
		Title:  a.Title,
		URL:    a.URL,
	}

	ValidateJSONs(t, a, m2)
}
