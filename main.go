package main

import (
	"flag"
)

var url URL
var method http_method
var methods []http_method

func main() {
	complexFlag := flag.Bool("complex", false, "Enable complex requests")
	flag.Parse()
	if *complexFlag {
		url, methods = complexArgumentParser()
		makeComplexRequest(url, methods)
	} else {
		url, method = simpleArgumentParser()
		makeRequest(url, method)
	}
}
