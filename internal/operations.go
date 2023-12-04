package internal

import (
	"bytes"
	"fmt"
	"net/http"
)

func printRequestId(r_id string) {
	if r_id != "" {
		fmt.Printf("x-request-id: %s \n", r_id)
	}
}
func printCookies(cookies []*http.Cookie) {
	if len(cookies) > 0 {
		fmt.Println("set-cookie: {")
		for _, cookie := range cookies {
			fmt.Printf("	\"%s\": \"%s\" \n", cookie.Name, cookie.Value)
		}
		fmt.Println("}")
	}
}
func printResponse(response *http.Response) {
	fmt.Printf("Status: %s \n", response.Status)
	fmt.Printf("HTTP: %s \n", response.Proto)
	fmt.Printf("Content-type: %s \n", response.Header.Get("Content-Type"))
	printRequestId(response.Header.Get("X-Request-ID"))
	printCookies(response.Cookies())
}
func requestUrl(url string, method string) (*http.Response, error) {
	request, err := http.NewRequest(string(method), string(url), bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println("Error al crear la solicitud "+method+":", err)
		panic("")
	}
	client := &http.Client{}
	return client.Do(request)
}
func MakeRequest(url string, method string) {
	res, err := requestUrl(url, method)

	if err != nil {
		fmt.Println("Error al hacer la solicitud "+method+":", err)
		panic("")
	}
	printResponse(res)
}
func makeComplexRequest(url string, methods []string) {
	fmt.Println("Complex request to " + url)
	for _, v := range methods {
		fmt.Printf("Method: %s \n", v)
		MakeRequest(url, v)
	}
}
