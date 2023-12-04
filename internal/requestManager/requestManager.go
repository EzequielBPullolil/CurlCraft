package requestmanager

import (
	argumentManager "github.com/EzequielK-source/CurlCraft/internal/argumentManager"
	basicRequest "github.com/EzequielK-source/CurlCraft/internal/basicRequest"
	complexRequest "github.com/EzequielK-source/CurlCraft/internal/complexRequest"
)

type requestManager struct {
	haveBodyData bool
	isComplex    bool
}

func RequestManager(haveBodyData bool, isComplex bool) requestManager {
	return requestManager{haveBodyData: haveBodyData, isComplex: isComplex}
}

func (r requestManager) MakeRequest(args []string) {
	url := argumentManager.Url(args)
	if r.isComplex {
		complexRequest.Request(url, argumentManager.Methods(args))
	} else {
		basicRequest.Request(url, argumentManager.Method(args))
	}
}
