package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

func printRequestId(r_id string) {
	if r_id != "" {
		color.Magenta("x-request-id: %s \n", r_id)
	}
}
func printCookies(cookies []*http.Cookie) {
	if len(cookies) > 0 {
		color.Blue("set-cookie: {")
		for _, cookie := range cookies {
			color.Blue("	\"%s\": \"%s\" \n", cookie.Name, cookie.Value)
		}
		color.Blue("}")
	}
}
func PrintResponse(response *http.Response, showBodyResponse bool) {
	printStatusCode(response.Status)
	color.Cyan("HTTP: %s \n", response.Proto)
	color.Magenta("Content-type: %s \n", response.Header.Get("Content-Type"))
	printRequestId(response.Header.Get("X-Request-ID"))
	printCookies(response.Cookies())
	if showBodyResponse {
		body, _ := io.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}
func printStatusCode(statusCode string) {
	if strings.Contains(statusCode, "2") {
		color.Cyan("Status: %s \n", statusCode)
	} else if strings.Contains(statusCode, "4") {
		color.Red("Status: %s \n", statusCode)
	} else {
		color.Blue("Status: %s \n", statusCode)
	}
}
