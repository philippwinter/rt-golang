package rt

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type Response struct {
	resp *http.Response
	t    *testing.T
}

func (r *Response) ExpectStatus(expected int) {
	if r.resp.StatusCode != expected {
		r.t.Fatalf("Status is %s, expected %s", r.resp.Status, Status(expected))
	}
}

func (r *Response) MustParseAsJson() map[string]interface{} {
	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		r.t.Fatalf("Could not read body into byte array: %+v", err)
	}

	return r.JsonToDocMap(&bytes)
}

func (r *Response) JsonToDocMap(bytes *[]byte) map[string]interface{} {
	var data interface{}

	err := json.Unmarshal(*bytes, &data)
	if err != nil {
		r.t.Fatalf("Could not unmarshal JSON: %+v", err)
	}

	return data.(map[string]interface{})
}
