package commands

import (
	"net/url"
	"testing"

	"github.com/Shaked/getpocket/utils"
)

type request struct {
	ret string
	err *utils.RequestError
}

func (r *request) Post(url string, values url.Values) ([]byte, *utils.RequestError) {
	return []byte(r.ret), r.err
}

func TestfixJSONArrayToObject(t *testing.T) {
	// This
	apiResult := []byte(`{"item_id":"1", "videos":[],"authors":[],"images":[]}`)
	expected := []byte(`{"item_id":"1", "videos":{},"authors":{},"images":{}}`)
	actual := fixJSONArrayToObject(apiResult)
	if string(expected) != string(actual) {
		t.Errorf("Actual value is worng %s", string(actual))
	}

}
