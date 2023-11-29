package internal

import (
	"os"
	"strings"
)

func appendMethod(a []http_method, e string) []http_method {
	if isHttpMethod(strings.ToUpper(e)) {
		return append(a, http_method(e))
	} else {
		return a
	}
}

func complexArgumentParser() (URL, []http_method) {
	var url URL
	var methods []http_method
	for _, v := range os.Args[1:] {
		methods = appendMethod(methods, v)
		if isURL(v) {
			url = URL(v)
		}
	}

	return url, methods
}

func simpleArgumentParser() (URL, http_method) {
	var method http_method
	var url URL
	for _, v := range os.Args[1:] {
		if isHttpMethod(strings.ToUpper(v)) {
			method = http_method(v)
			continue
		}
		if isURL(v) {
			url = URL(v)
			continue
		}
	}

	return url, method
}
