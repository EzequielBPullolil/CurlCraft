package argumentmanager

import (
	"bytes"
	"io"

	"github.com/EzequielK-source/CurlCraft/internal"
)

func ManageArguments(args []string) (string, []string, string, io.Reader) {
	var contentType string
	var methods []string

	for _, argument := range args[1:] {
		if internal.IsContentType(argument) {
			contentType = argument
			continue
		}

		if internal.IsHttpMethod(argument) {
			methods = append(methods, argument)
			continue
		}
	}
	return args[0], methods, contentType, bytes.NewReader([]byte(""))
}
