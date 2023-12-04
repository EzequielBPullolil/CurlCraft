package complexrequest

import (
	"github.com/EzequielK-source/CurlCraft/internal"
)

func Request(url string, methods []string) {
	for _, m := range methods {
		internal.MakeRequest(url, m)
	}
}
