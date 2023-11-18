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
	fmt.Println("El url es " + url)

	if len(methods) == 0 {
		fmt.Println("Metodo " + method)
	} else {
		fmt.Println("Metodos")
		for _, v := range methods {
			fmt.Println(v)
		}
	}
}
