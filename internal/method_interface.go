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

func IsHttpMethod(e string) bool {
	e = strings.ToUpper(e)
	return e == POST || e == GET || e == DELETE || e == PUT || e == PATCH
}

func IsContentType(e string) bool {
	e = strings.ToUpper(e)

	return e == "JSON" || e == "FORM" || e == "PLAIN" || e == "XML" || e == "HTML" || e == "XFORM"
}
