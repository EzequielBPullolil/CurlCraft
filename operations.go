package main

import "net/http"

func makeRequest(u URL, m http_method) (string, error) {
	var status string
	res, err := http.Get(string(u))
	if err != nil {
		return "", err
	}
	status = res.Status

	return status, nil
}
