package modify

import (
	"encoding/json"
	"testing"
)

func ValidateJSONs(t *testing.T, a Action, m interface{}) {
	j, e := json.Marshal(a)
	if nil != e {
		t.Error(e)
	}
	expectedJSON, _ := json.Marshal(m)
	if string(j) != string(expectedJSON) {
		t.Errorf("Expected %s is not the same as actual %s", expectedJSON, j)
	}
}
