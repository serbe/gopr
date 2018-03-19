package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", cfg.Web.CORS_URL)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With",
		)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<p>RemoteAddr: %s</p>", r.RemoteAddr)
	errChkMsg("checkHandler fmt.Fprintf", err)
	for _, header := range headers {
		str := r.Header.Get(header)
		if str == "" {
			continue
		}
		_, err = fmt.Fprintf(w, "<p>%s: %s</p>", header, str)
		errChkMsg("checkHandler fmt.Fprintf", err)
	}
}

func getProxy(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Title string    `json:"title"`
		Proxy adb.Proxy `json:"proxy"`
	}
	id := toInt(chi.URLParam(r, "id"))
	proxy, _ := db.ProxyGetByID(id)
	ctx := context{Title: "Proxy", Proxy: proxy}
	render.DefaultResponder(w, r, ctx)
}

func listProxies(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	proxies, _ := db.ProxyGetAll()
	ctx := context{Title: "List all proxies", Proxies: proxies}
	render.DefaultResponder(w, r, ctx)
}

func listWorkProxies(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	proxies, _ := db.ProxyGetAllWorking()
	ctx := context{Title: "List working proxies", Proxies: proxies}
	render.DefaultResponder(w, r, ctx)
}
func listAnonProxies(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	proxies, _ := db.ProxyGetAllAnonymous()
	ctx := context{Title: "List anonymous proxies", Proxies: proxies}
	render.DefaultResponder(w, r, ctx)
}

func getCounts(w http.ResponseWriter, r *http.Request) {
	type context struct {
		Title string `json:"title"`
		All   int64  `json:"all"`
		Work  int64  `json:"work"`
		Anon  int64  `json:"anon"`
	}
	all := db.ProxyGetAllCount()
	work := db.ProxyGetAllWorkCount()
	anon := db.ProxyGetAllAnonymousCount()
	ctx := context{Title: "Proxies counts", All: all, Work: work, Anon: anon}
	render.DefaultResponder(w, r, ctx)
}
