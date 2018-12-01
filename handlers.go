package main

import (
	"fmt"
	"net/http"
)

var headers = []string{
	"HTTP_VIA",
	"HTTP_X_FORWARDED_FOR",
	"HTTP_FORWARDED_FOR",
	"HTTP_X_FORWARDED",
	"HTTP_FORWARDED",
	"HTTP_CLIENT_IP",
	"HTTP_FORWARDED_FOR_IP",
	"VIA",
	"X_FORWARDED_FOR",
	"FORWARDED_FOR",
	"X_FORWARDED",
	"FORWARDED",
	"CLIENT_IP",
	"FORWARDED_FOR_IP",
	"HTTP_PROXY_CONNECTION",
}

func checkHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "<p>RemoteAddr: %s</p>", req.RemoteAddr)
	errChkMsg("checkHandler fmt.Fprintf", err)
	for _, header := range headers {
		str := req.Header.Get(header)
		if str == "" {
			continue
		}
		_, err = fmt.Fprintf(w, "<p>%s: %s</p>", header, str)
		errChkMsg("checkHandler fmt.Fprintf", err)
	}
}
