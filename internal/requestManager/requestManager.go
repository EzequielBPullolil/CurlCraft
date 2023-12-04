package requestmanager

import (
	"github.com/EzequielK-source/CurlCraft/internal"
	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
)

type requestManager struct {
	haveBodyData bool
	isComplex    bool
	url          string
	methods      []string
	contentType  string
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
	return requestManager{
		haveBodyData: haveBodyData,
		isComplex:    isComplex,
		url:          argumentManager.Url(args),
		methods:      argumentManager.Methods(args),
		contentType:  argumentManager.ContentType(args),
	}
}

func (r requestManager) MakeRequest() {
	if r.isComplex {
		r.complexRequest()
	} else {
		r.basicRequest()
	}
}
