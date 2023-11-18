package main

import "fmt"

func main() {
	url, methods := argumentManager()

	fmt.Println("Request to " + url)
	for _, v := range methods {
		fmt.Println("Request to " + v)

	}
}
