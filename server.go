package main

import (
	"github.com/valyala/fasthttp"
)

func initServer() {
	mux := func(ctx *fasthttp.RequestCtx) {
		if ctx.IsGet() {
			switch string(ctx.Path()) {
			// case "/loaderio-a756090c2ec9b33ba21d957b28485477.txt":
			// 	fmt.Fprintf(ctx, "loaderio-a756090c2ec9b33ba21d957b28485477")
			case "/check":
				checkHandler(ctx)
			case "/api/proxies/all":
				if checkAuth(ctx) {
					listProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case "/api/proxies/work":
				if checkAuth(ctx) {
					listWorkProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case "/api/proxies/anon":
				if checkAuth(ctx) {
					listAnonProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case "/api/proxies/counts":
				if checkAuth(ctx) {
					getCounts(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		} else if ctx.IsPost() {
			switch string(ctx.Path()) {
			case "/login":
				login(ctx)
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		} else {
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	err := fasthttp.ListenAndServe(":"+cfg.Web.Port, mux)
	errChkMsg("ListenAndServe", err)
}
