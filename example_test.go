package rt

import (
	"github.com/philippwinter/rt-golang/testserver"
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
		t.Fatalf("Did not receive expected input: %+v", docMap)
	}
}
