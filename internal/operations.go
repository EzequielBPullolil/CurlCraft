package internal

import (
	"bytes"
	"fmt"
	"net/http"
)

func printResponse(response *http.Response) {
	cookies := response.Cookies()
	request_id := response.Header.Get("X-Request-ID")
	fmt.Printf("Status: %s \n", response.Status)
	fmt.Printf("HTTP: %s \n", response.Proto)
	fmt.Printf("Content-type: %s \n", response.Header.Get("Content-Type"))
	if request_id != "" {
		fmt.Printf("x-request-id: %s \n", request_id)
	}
	if len(cookies) > 0 {
		fmt.Println("set-cookie: {")
		for _, cookie := range cookies {
			fmt.Printf("	\"%s\": \"%s\" \n", cookie.Name, cookie.Value)
		}
		fmt.Println("}")
	}
}
func requestUrl(u string, m string) (*http.Response, error) {
	request, err := http.NewRequest(string(m), string(u), bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println("Error al crear la solicitud "+m+":", err)
		panic("")
	}
	client := &http.Client{}
	return client.Do(request)
}
func MakeRequest(u string, m string) {
	res, err := requestUrl(u, m)

	if err != nil {
		fmt.Println("Error al hacer la solicitud "+m+":", err)
		panic("")
	}
	printResponse(res)
}
func makeComplexRequest(u string, ms []string) {
	fmt.Println("Complex request to " + u)
	for _, v := range ms {
		fmt.Printf("Method: %s \n", v)
		MakeRequest(u, v)
	}
}
