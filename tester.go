package rt

import (
	"log"
	"net/http"
	"testing"
)

type Tester struct {
	testing.T
}

func (tester *Tester) Must(resp *http.Response, err error) *Response {
	if err != nil {
		log.Printf("Could not perform request: %+v\n", err)
		tester.FailNow()
	}

	return &Response{
		resp,
		&tester.T,
	}
}
