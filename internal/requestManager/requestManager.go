package requestmanager

import (
	"bytes"
	"io"
	"net/http"

	"github.com/EzequielK-source/CurlCraft/internal"
	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
)

type requestManager struct {
	isComplex   bool
	url         string
	methods     []string
	contentType string
	client      http.Client
	bodyData    io.Reader
}
}

func (r requestManager) basicRequest() {
	method := "get"
	if len(r.methods) > 0 {
		method = r.methods[len(r.methods)-1]
	}
	internal.MakeRequest(r.url, method)
}

func (r requestManager) complexRequest() {
	for _, m := range r.methods {
		internal.MakeRequest(r.url, m)
	}
}

func RequestManager(haveBodyData bool, isComplex bool, args []string) requestManager {
	url, methods, contentType := argumentManager.ManageArguments(haveBodyData, args)
	return requestManager{
		isComplex:   isComplex,
		url:         url,
		methods:     methods,
		contentType: contentType,
		client:      http.Client{},
		bodyData:    bytes.NewBuffer(nil),
	}
}

func (r requestManager) MakeRequest() {
	if r.isComplex {
		r.complexRequest()
	} else {
		r.basicRequest()
	}
}
