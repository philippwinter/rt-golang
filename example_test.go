package rt

import (
	"github.com/philippwinter/rt-golang/testserver"
	"log"
	"testing"
)

var testServer = &Server{
	Protocol: "http",
	Host:     testserver.HOST,
	Port:     testserver.PORT,
}

func TestSimpleGet(t *testing.T) {
	tester, client := Init(t, InsecureTLSConfig())

	resp := tester.Must(client.Get(testServer.URL("/")))
	docMap := resp.MustParseAsJson()

	if docMap["hello"] != "world" {
		log.Printf("Did not receive expected input: %+v\n", docMap)
		t.Error()
	}
}
