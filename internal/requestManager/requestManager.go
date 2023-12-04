package requestmanager

import (
	"github.com/EzequielK-source/CurlCraft/internal"
	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
	complexRequest "github.com/EzequielK-source/CurlCraft/internal/complexRequest"
)

type requestManager struct {
	haveBodyData bool
	isComplex    bool
	url          string
	methods      []string
}

func (r requestManager) basicRequest() {
	method := "get"
	if len(r.methods) > 0 {
		method = r.methods[len(r.methods)-1]
	}
	internal.MakeRequest(r.url, method)
}

func RequestManager(haveBodyData bool, isComplex bool, args []string) requestManager {
	return requestManager{haveBodyData: haveBodyData, isComplex: isComplex, url: argumentManager.Url(args), methods: argumentManager.Methods(args)}
}

func (r requestManager) MakeRequest() {
	if r.isComplex {
		complexRequest.Request(r.url, r.methods)
	} else {
		r.basicRequest()
	}
}
