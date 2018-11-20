package main

import (
	"bytes"

	"github.com/valyala/fasthttp"
)

var (
	strPOST            = []byte("POST")
	strGET             = []byte("GET")
	pathCheck          = []byte("/check")
	pathLogin          = []byte("/login")
	pathAPIAllProxy    = []byte("/api/proxies/all")
	pathAPIWorkProxy   = []byte("/api/proxies/work")
	pathAPIAnonProxy   = []byte("/api/proxies/anon")
	pathAPICountsProxy = []byte("/api/proxies/counts")
)

func initServer() {
	mux := func(ctx *fasthttp.RequestCtx) {
		switch {
		case bytes.Equal(ctx.Method(), strGET):
			switch {
			case bytes.Equal(ctx.Path(), pathCheck):
				checkHandler(ctx)
			case bytes.Equal(ctx.Path(), pathAPIAllProxy):
				if checkAuth(ctx) {
					listProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case bytes.Equal(ctx.Path(), pathAPIWorkProxy):
				if checkAuth(ctx) {
					listWorkProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case bytes.Equal(ctx.Path(), pathAPIAnonProxy):
				if checkAuth(ctx) {
					listAnonProxies(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			case bytes.Equal(ctx.Path(), pathAPICountsProxy):
				if checkAuth(ctx) {
					getCounts(ctx)
				} else {
					ctx.Error("Not Authorized", fasthttp.StatusUnauthorized)
				}
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		case bytes.Equal(ctx.Method(), strPOST):
			switch {
			case bytes.Equal(ctx.Path(), pathLogin):
				login(ctx)
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	err := fasthttp.ListenAndServe(":"+cfg.Web.Port, mux)
	errChkMsg("ListenAndServe", err)
}
