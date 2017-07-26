package rt

import (
	"fmt"
	"net/http"
)

// Important MIME types
const (
	JSON = "application/json"
)

// Important HTTP methods
const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	HEAD   = "HEAD"
	DELETE = "DELETE"
)

func Status(statusCode int) string {
	return fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode))
}
