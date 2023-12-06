package consolewriter

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
func WriteResponse(response *http.Response, showBodyResponse bool) {
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
	var colored string
	if strings.Contains(statusCode, "3") {
		colored = RedirectSp(statusCode)
	} else if strings.Contains(statusCode, "4") {
		colored = FailSp(statusCode)
	} else if strings.Contains(statusCode, "5") {
		colored = ServerErrorSp(statusCode)
	} else {
		colored = SuccesSp(statusCode)
	}
	color.Magenta("Satuts: %s", colored)
}
