package basicrequest

import (
	"fmt"

	"github.com/EzequielK-source/CurlCraft/internal"
)

func Request(url string, method string) {
	internal.MakeRequest(url, method)
	fmt.Println(url)
	fmt.Println(method)
}
