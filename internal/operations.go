package internal

import (
	"fmt"
	"io/ioutil"
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
func PrintResponse(response *http.Response) {
	fmt.Printf("Status: %s \n", response.Status)
	fmt.Printf("HTTP: %s \n", response.Proto)
	fmt.Printf("Content-type: %s \n", response.Header.Get("Content-Type"))
	printRequestId(response.Header.Get("X-Request-ID"))
	printCookies(response.Cookies())
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
