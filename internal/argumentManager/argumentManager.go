package argumentmanager

import (
	"github.com/EzequielK-source/CurlCraft/internal"
)

func ManageArguments(haveBodyData bool, args []string) (string, []string, string) {
	var contentType string
	var methods []string

	for _, argument := range args {
		if haveBodyData {
			if internal.IsContentType(argument) {
				contentType = argument
			}

			if internal.IsHttpMethod(argument) {
				methods = append(methods, argument)
			}
		}
	}

	if len(methods) < 1 {
		methods = append(methods, "GET")
	}
	return args[0], methods, contentType
}
