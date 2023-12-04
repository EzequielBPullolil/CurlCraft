package complexrequest

import (
	"fmt"

	"github.com/EzequielK-source/CurlCraft/internal"
)

func Request(url string, methods []string) {
	for _, m := range methods {
		fmt.Println("Method: " + m)
		internal.MakeRequest(url, m)
	}
}
