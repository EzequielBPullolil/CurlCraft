package requestmanager

import (
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

func (r requestManager) request(method string) {
	request, err := http.NewRequest(method, r.url, r.bodyData)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := r.client.Do(request)

	if err != nil {
		panic(err)
	}

	internal.PrintResponse(res)
}

func (r requestManager) basicRequest() {
	method := "get"
	if len(r.methods) > 0 {
		method = r.methods[len(r.methods)-1]
	}
	r.request(method)
}

func (r requestManager) complexRequest() {
	for _, m := range r.methods {
		r.request(m)
	}
}

func RequestManager(body string, isComplex bool, args []string) requestManager {
	url, methods, contentType := argumentManager.ManageArguments(args)
	bodyData := argumentManager.ContentTypeEncoder(contentType, body)
	return requestManager{
		isComplex:   isComplex,
		url:         url,
		methods:     methods,
		contentType: contentType,
		client:      http.Client{},
		bodyData:    bodyData,
	}
}

func (r requestManager) MakeRequest() {
	if r.isComplex {
		r.complexRequest()
	} else {
		r.basicRequest()
	}

	r.client.CloseIdleConnections()
}
