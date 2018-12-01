package main

import (
	"net/http"
)

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			switch req.URL.Path {
			case "/check":
				checkHandler(w, req)
			default:
				http.Error(w, "not found", http.StatusNotFound)
			}
		default:
			http.Error(w, "not found", http.StatusNotFound)
		}
	})

	err := http.ListenAndServe(":"+cfg.Web.Port, mux)
	errChkMsg("ListenAndServe", err)
}
