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
