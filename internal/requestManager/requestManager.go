package requestmanager

import (
	"io"
	"net/http"

	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
	consoleWriter "github.com/EzequielK-source/CurlCraft/internal/consoleWriter"
	"github.com/fatih/color"
)

type requestManager struct {
	isComplex        bool
	showBodyResponse bool
	url              string
	methods          []string
	contentType      string
	client           http.Client
	bodyData         io.Reader
}

func (r requestManager) request(method string) {
	request, err := http.NewRequest(method, r.url, r.bodyData)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", r.contentType)

	color.Green("Method: " + method)
	res, err := r.client.Do(request)

	if err != nil {
		panic(err)
	}

	consoleWriter.PrintResponse(res, r.showBodyResponse)
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

func RequestManager(body string, isComplex bool, showBodyResponse bool, args []string) requestManager {
	url, methods, contentType := argumentManager.ManageArguments(args)
	bodyData := argumentManager.ContentTypeEncoder(contentType, body)
	return requestManager{
		isComplex:        isComplex,
		url:              url,
		methods:          methods,
		contentType:      contentType,
		client:           http.Client{},
		bodyData:         bodyData,
		showBodyResponse: showBodyResponse,
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
