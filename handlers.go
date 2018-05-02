package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serbe/adb"
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

func listProxies(w http.ResponseWriter, req *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(w, req)
	proxies, _ := db.ProxyGetAll()
	err := json.NewEncoder(w).Encode(context{Title: "List all proxies", Proxies: proxies})
	errChkMsg("listProxies Encode", err)
}

func listWorkProxies(w http.ResponseWriter, req *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(w, req)
	proxies, _ := db.ProxyGetAllWorking()
	err := json.NewEncoder(w).Encode(context{Title: "List working proxies", Proxies: proxies})
	errChkMsg("listWorkProxies Encode", err)
}

func listAnonProxies(w http.ResponseWriter, req *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(w, req)
	proxies, _ := db.ProxyGetAllAnonymous()
	err := json.NewEncoder(w).Encode(context{Title: "List anonymous proxies", Proxies: proxies})
	errChkMsg("listWorkProxies Encode", err)
}

func getCounts(w http.ResponseWriter, req *http.Request) {
	type context struct {
		Title string `json:"title"`
		All   int64  `json:"all"`
		Work  int64  `json:"work"`
		Anon  int64  `json:"anon"`
	}
	cors(w, req)
	all := db.ProxyGetAllCount()
	work := db.ProxyGetAllWorkCount()
	anon := db.ProxyGetAllAnonymousCount()
	err := json.NewEncoder(w).Encode(context{Title: "Proxies counts", All: all, Work: work, Anon: anon})
	errChkMsg("getCounts Encode", err)
}

func cors(w http.ResponseWriter, req *http.Request) {
	if cfg.Web.CORS {
		w.Header().Set("Access-Control-Allow-Origin", cfg.Web.CorsURL)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With",
		)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

// func getProxy(w http.ResponseWriter, r *http.Request) {
// 	type context struct {
// 		Title string    `json:"title"`
// 		Proxy adb.Proxy `json:"proxy"`
// 	}
// 	id := toInt(chi.URLParam(r, "id"))
// 	proxy, _ := db.ProxyGetByID(id)
// 	ctx := context{Title: "Proxy", Proxy: proxy}
// 	render.DefaultResponder(w, r, ctx)
// }
