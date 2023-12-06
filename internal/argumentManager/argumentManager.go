package argumentmanager

import (
	"bytes"
	"io"

	"github.com/EzequielK-source/CurlCraft/internal"
)

func ManageArguments(args []string) (string, []string, string) {
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
	return args[0], methods, contentType
}
func ContentTypeEncoder(contentType string, data string) io.Reader {
	if !internal.IsContentType(contentType) {
		panic("Not valid content type")
	}

	return encondedContentType(contentType, data)
}
}
