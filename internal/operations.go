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
func MakeRequest(url, method string) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(nil))
	if err != nil {
		return
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return
	}
	printResponse(res)
}
