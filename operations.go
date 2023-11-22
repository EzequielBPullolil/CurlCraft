package main

import (
	"fmt"
	"net/http"
)

func printResponse(response *http.Response) {
	fmt.Printf("Status: %s \n", response.Status)
	fmt.Printf("HTTP: %s \n", response.Proto)
	fmt.Printf("Content-type: %s \n", response.Header.Get("Content-Type"))
	fmt.Printf("x-request-id: %s \n", response.Header.Get("X-Request-ID"))
	fmt.Println("set-cookie: {")
	for _, cookie := range response.Cookies() {
		fmt.Printf("	\"%s\": \"%s\" \n", cookie.Name, cookie.Value)
	}
	fmt.Println("}")
}

func makeRequest(u URL, m http_method) {
	res, err := http.Get(string(u))
	if err != nil {
		panic("Failed to request url " + u + "fix our entry")
	}
	printResponse(res)
}
}
