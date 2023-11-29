package internal

import (
	"strings"
)

type http_method string
type URL string

const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
	PUT    = "PUT"
	PATCH  = "PATCH"
)

func isHttpMethod(e string) bool {
	e = strings.ToUpper(e)
	return e == POST || e == GET || e == DELETE || e == PUT || e == PATCH
}

func isURL(e string) bool {
	return strings.HasPrefix(e, "http")
}
