package rt

import (
	"fmt"
	"net/http"
)

type Method int
type Mime int

// Important MIME types
const (
	JSON Mime = iota
)

// Important HTTP methods
const (
	GET Method = iota
	POST
	PATCH
	HEAD
	DELETE
)

func Status(statusCode int) string {
	return fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode))
}

func (method Method) String() string {
	switch method {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PATCH:
		return "PATCH"
	case HEAD:
		return "HEAD"
	case DELETE:
		return "DELETE"
	default:
		panic("Undefined method")
	}
}

func (m Mime) String() string {
	switch m {
	case JSON:
		return "application/json"
	default:
		panic("Undefined mime")
	}
}
