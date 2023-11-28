package main

import (
	"flag"
)

var url URL
var method http_method
var methods []http_method

func main() {
	isComplex := flag.Bool("complex", false, "Enable complex requests")
	isHtmlResponseAllowed := flag.Bool("h", false, "Enable complex html response")
	flag.Parse()

	manageRequest(*isComplex, *isHtmlResponseAllowed, url)
}
