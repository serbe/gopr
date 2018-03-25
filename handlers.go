package main

import (
	"encoding/json"
	"fmt"

	"github.com/serbe/adb"
	"github.com/valyala/fasthttp"
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

func checkHandler(ctx *fasthttp.RequestCtx) {
	_, err := fmt.Fprintf(ctx, "<p>RemoteAddr: %s</p>", ctx.RemoteAddr())
	errChkMsg("checkHandler fmt.Fprintf", err)
	for _, header := range headers {
		str := ctx.Request.Header.Peek(header)
		if str == nil {
			continue
		}
		_, err = fmt.Fprintf(ctx, "<p>%s: %s</p>", header, str)
		errChkMsg("checkHandler fmt.Fprintf", err)
	}
}

func listProxies(ctx *fasthttp.RequestCtx) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(ctx)
	proxies, _ := db.ProxyGetAll()
	err := json.NewEncoder(ctx).Encode(context{Title: "List all proxies", Proxies: proxies})
	errChkMsg("listProxies Encode", err)
}

func listWorkProxies(ctx *fasthttp.RequestCtx) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(ctx)
	proxies, _ := db.ProxyGetAllWorking()
	err := json.NewEncoder(ctx).Encode(context{Title: "List working proxies", Proxies: proxies})
	errChkMsg("listWorkProxies Encode", err)
}

func listAnonProxies(ctx *fasthttp.RequestCtx) {
	type context struct {
		Title   string      `json:"title"`
		Proxies []adb.Proxy `json:"proxies"`
	}
	cors(ctx)
	proxies, _ := db.ProxyGetAllAnonymous()
	err := json.NewEncoder(ctx).Encode(context{Title: "List anonymous proxies", Proxies: proxies})
	errChkMsg("listWorkProxies Encode", err)
}

func getCounts(ctx *fasthttp.RequestCtx) {
	type context struct {
		Title string `json:"title"`
		All   int64  `json:"all"`
		Work  int64  `json:"work"`
		Anon  int64  `json:"anon"`
	}
	cors(ctx)
	all := db.ProxyGetAllCount()
	work := db.ProxyGetAllWorkCount()
	anon := db.ProxyGetAllAnonymousCount()
	err := json.NewEncoder(ctx).Encode(context{Title: "Proxies counts", All: all, Work: work, Anon: anon})
	errChkMsg("getCounts Encode", err)
}

func cors(ctx *fasthttp.RequestCtx) {
	if cfg.Web.CORS {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", cfg.Web.CORS_URL)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", cfg.Web.CORS_URL)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		ctx.Response.Header.Set("Access-Control-Max-Age", "3600")
		ctx.Response.Header.Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With",
		)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
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
