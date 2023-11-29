package basicrequest

import (
	"github.com/EzequielK-source/CurlCraft/internal"
)

func Request(url string, method string) {
	internal.MakeRequest(url, method)
}
