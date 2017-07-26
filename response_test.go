package rt

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestResponse_JsonToDocMap(t *testing.T) {
	readCloser := ioutil.NopCloser(bytes.NewBufferString(`{"name": "rt"}`))

	resp := &Response{
		&http.Response{
			StatusCode: http.StatusOK,
			Status:     http.StatusText(http.StatusOK),
			Body:       readCloser,
		},
		t,
	}

	docMap := resp.MustParseAsJson()
	if docMap["name"] != "rt" {
		log.Printf("JSON wrongfully unwrapped\n")
		t.Fatal()
	}
}
