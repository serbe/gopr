package main

import (
	"net/http"
)

func initServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			switch req.URL.Path {
			case "/login":
				login(w, req)
			default:
				http.Error(w, "not found", http.StatusNotFound)
			}
		case http.MethodGet:
			switch req.URL.Path {
			// case "/loaderio-a756090c2ec9b33ba21d957b28485477.txt":
			// 	fmt.Fprintf(ctx, "loaderio-a756090c2ec9b33ba21d957b28485477")
			case "/check":
				checkHandler(w, req)
			case "/api/proxies/all":
				if checkAuth(w, req) {
					listProxies(w, req)
				} else {
					http.Error(w, "Not Authorized", http.StatusUnauthorized)
				}
			case "/api/proxies/work":
				if checkAuth(w, req) {
					listWorkProxies(w, req)
				} else {
					http.Error(w, "Not Authorized", http.StatusUnauthorized)
				}
			case "/api/proxies/anon":
				if checkAuth(w, req) {
					listAnonProxies(w, req)
				} else {
					http.Error(w, "Not Authorized", http.StatusUnauthorized)
				}
			case "/api/proxies/counts":
				if checkAuth(w, req) {
					getCounts(w, req)
				} else {
					http.Error(w, "Not Authorized", http.StatusUnauthorized)
				}
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
