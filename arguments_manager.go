package main

import (
	"os"
	"strings"
)

func appendMethod(a []http_method, e string) []http_method {
	return append(a, http_method(e))
}

func complexArgumentParser() (URL, []http_method) {
	var methods []http_method
	for _, v := range os.Args[1:] {
		methods = appendMethod(methods, v)
	}

	return "", methods
}

func simpleArgumentParser() (URL, http_method) {
	var method http_method
	var url URL
	for _, v := range os.Args[1:] {
		if isHttpMethod(strings.ToUpper(v)) {
			method = http_method(v)
		} else {
			url = URL(v)
		}
	}

	return url, method
}
