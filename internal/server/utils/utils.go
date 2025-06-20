package utils

import "net/http"

func CopyHeaders(req, proxyReq http.Header) {
	for key, value := range req {
		for _, header := range value {
			proxyReq.Add(key, header)
		}
	}
}
