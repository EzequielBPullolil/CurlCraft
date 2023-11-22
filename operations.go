package main

import (
	"bytes"
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
func requestUrl(u URL, m http_method) (*http.Response, error) {
	request, err := http.NewRequest(string(m), string(u), bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println("Error al crear la solicitud "+m+":", err)
		panic("")
	}
	client := &http.Client{}
	return client.Do(request)
}
func makeRequest(u URL, m http_method) {
	res, err := requestUrl(u, m)

	if err != nil {
		fmt.Println("Error al hacer la solicitud "+m+":", err)
		panic("")
	}
	printResponse(res)
}
func makeComplexRequest(u URL, ms []http_method) {
	fmt.Println("Complex request to " + u)
	for _, v := range ms {
		fmt.Printf("Method: %s \n", v)
		makeRequest(u, v)
	}
}
