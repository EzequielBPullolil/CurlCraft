package main

import (
	"os"
)

func appendMethod(a []http_method, e string) []http_method {
	return append(a, http_method(e))
}

func argumentManager() (url, []http_method) {
	// var url string
	var methods []http_method
	for _, v := range os.Args[1:] {
		methods = appendMethod(methods, v)
	}

	return "", methods
}
