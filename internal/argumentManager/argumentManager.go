package argumentmanager

import (
	internal "github.com/EzequielK-source/CurlCraft/internal"
)

func Url(args []string) string {
	if len(args) < 0 {
		panic("CurlCraft needs a url or parameter")
	}

	return args[0]
}

func Methods(args []string) []string {
	var methods []string
	for _, v := range args[1:] {
		if internal.IsHttpMethod(v) {
			methods = append(methods, v)
		}
	}
	return methods
}

func Method(args []string) string {
	for _, v := range args[1:] {
		if internal.IsHttpMethod(v) {
			return v
		}
	}

	return "GET"
}
