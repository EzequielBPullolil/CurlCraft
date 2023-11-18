package main

import (
	"flag"
	"fmt"
)

func main() {
	var url URL
	var method http_method
	var methods []http_method
	complexFlag := flag.Bool("complex", false, "Allows complex requests")

	flag.Parse()
	if *complexFlag {
		url, methods = complexArgumentParser()
	} else {
		url, method = simpleArgumentParser()
	}
	if len(methods) == 0 {
		resp, err := makeRequest(url, method)
		if err == nil {
			fmt.Print(resp)
		}
	} else {
		for _, m := range methods {
			resp, err := makeRequest(url, m)
			if err == nil {
				fmt.Print(resp)
			}
		}
	}
}
