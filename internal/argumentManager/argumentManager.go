package argumentmanager

import (
	"bytes"
	"io"

	"github.com/EzequielK-source/CurlCraft/internal"
)

func manageJsonArgument(args []string) string {
	var s string
	for _, v := range args {
		s = s + v
	}
	return s
}

func ManageArguments(haveBodyData bool, args []string) (string, []string, string, io.Reader) {
	var contentType string
	var methods []string
	var bodyData string

	for i, argument := range args[1:] {
		if internal.IsContentType(argument) {
			contentType = argument
			continue
		}

		if internal.IsHttpMethod(argument) {
			methods = append(methods, argument)
			continue
		}
		if haveBodyData {
			bodyData = manageJsonArgument(args[i+1:])
			break
		}
	}
	return args[0], methods, contentType, bytes.NewReader([]byte(bodyData))
}
