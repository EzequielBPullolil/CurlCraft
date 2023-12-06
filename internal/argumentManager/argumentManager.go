package argumentmanager

import (
	"bytes"
	"fmt"
	"io"
	"strings"

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
	return args[0], methods, formatContentType(contentType)
}

func formatContentType(content string) string {
	content = strings.ToUpper(content)
	if content == "HTML" {
		return "text/html"
	} else if content == "XML" {
		return "application/xml"
	} else if content == "XFORM" {
		return "application/x-www-form-urlencoded"
	} else if content == "FORM" {
		return "application/form-data"
	} else {
		return "application/json"
	}
}
func ContentTypeEncoder(contentType string, data string) io.Reader {
	return encondedContentType(contentType, data)
}

func encondedContentType(contentType string, data string) io.Reader {
	if contentType == "XML" {
		return xml(data)
	} else if contentType == "JSON" {
		return json(data)
	} else {
		return simpleKeyValue(data)
	}
}

func json(data string) io.Reader {
	return bytes.NewBuffer([]byte(data))
}
func simpleKeyValue(data string) io.Reader {
	var finalData string
	for _, v := range data {
		if v == ' ' {
			continue
		}
		if v == ',' {
			finalData = finalData + "&"
		} else if v == ':' {
			finalData = finalData + "="
		} else if v == '"' {
			continue
		} else if v != '{' && v != '}' {
			finalData = finalData + string(v)
		}
	}

	return bytes.NewBuffer([]byte(finalData))
}

func xml(data string) io.Reader {
	finalData := "<data>"
	for i := 0; i < len(data); i++ {
		if data[i] == '"' {
			finalData = finalData + crearEtiqueta(data, &i)
		}
	}
	finalData = finalData + "</data>"
	return bytes.NewBuffer([]byte(finalData))
}

func crearEtiqueta(data string, indice *int) string {
	var key string
	var value string
	cantidadDeComs := 4
	for cantidadDeComs > 0 {
		if data[*indice] != ' ' && data[*indice] != ':' {

			if data[*indice] == '"' {
				cantidadDeComs--
			} else {
				if cantidadDeComs > 2 {
					key = key + string(data[*indice])
				} else {
					value = value + string(data[*indice])
				}
			}
		}
		*indice++
	}
	return fmt.Sprintf("<%s>%s</%s>", key, value, key)
}
