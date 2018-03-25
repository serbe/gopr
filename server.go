package main

import (
	"github.com/valyala/fasthttp"
)

func initServer() {
	m := func(ctx *fasthttp.RequestCtx) {
		if ctx.IsPost() {
			switch string(ctx.Path()) {
			case "/login":
				login(ctx)
			default:
				ctx.Error("not found", fasthttp.StatusNotFound)
			}
		} else if ctx.IsGet() {
			switch string(ctx.Path()) {
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
		} else {
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	err := fasthttp.ListenAndServe(":"+cfg.Web.Port, m)
	errChkMsg("ListenAndServe", err)
}

// package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-chi/jwtauth"
// 	"github.com/go-chi/render"
// )

// func initServer() {
// 	tokenAuth = jwtauth.New("HS256", sKey, nil)

// 	r := chi.NewRouter()

// 	if cfg.Web.Log {
// 		r.Use(middleware.Logger)
// 	}
// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.Recoverer)
// 	r.Use(middleware.Timeout(60 * time.Second))
// 	if cfg.Web.CORS {
// 		r.Use(corsHandler)
// 	}

// 	// Check
// 	r.Get("/check", checkHandler)

// 	// Auth
// 	r.Group(func(r chi.Router) {
// 		r.Post("/login", login)
// 	})

// 	// REST API
// 	r.Group(func(r chi.Router) {
// 		if cfg.Web.Auth {
// 			r.Use(jwtauth.Verifier(tokenAuth))
// 			r.Use(jwtauth.Authenticator)
// 		}

// 		r.Use(render.SetContentType(render.ContentTypeJSON))

// 		r.Route("/api/proxies", func(r chi.Router) {
// 			r.Get("/{id}", getProxy)
// 			r.Get("/all", listProxies)
// 			r.Get("/work", listWorkProxies)
// 			r.Get("/anon", listAnonProxies)
// 			r.Get("/counts", getCounts)
// 		})
// 	})

// 	err := http.ListenAndServe(":"+cfg.Web.Port, r)
// 	errChkMsg("ListenAndServe", err)
// }
