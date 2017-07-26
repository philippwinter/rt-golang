package rt

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type Response struct {
	resp *http.Response
	t    *testing.T
}

func (r *Response) ExpectStatus(expected int) {
	if r.resp.StatusCode != expected {
		log.Printf("Status is %s, expected %s\n", r.resp.Status, Status(expected))
		r.t.Error()
	}
}

func (r *Response) MustParseAsJson() map[string]interface{} {
	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		log.Printf("Could not read body into byte array: %+v\n", err)
		r.t.Error()
	}

	return r.JsonToDocMap(&bytes)
}

func (r *Response) JsonToDocMap(bytes *[]byte) map[string]interface{} {
	var data interface{}

	err := json.Unmarshal(*bytes, &data)
	if err != nil {
		log.Printf("Could not unmarshal JSON: %+v\n", err)
		r.t.Error()
	}

	return data.(map[string]interface{})
}
