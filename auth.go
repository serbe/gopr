package main

import (
	"bytes"
	"encoding/json"

	jwtGo "github.com/dgrijalva/jwt-go"

	"github.com/valyala/fasthttp"
)

var sKey = []byte("aH0tH3P5up3RdYP3r53crEt")

type loginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jsonData struct {
	Admin bool   `json:"admin"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func login(ctx *fasthttp.RequestCtx) {
	var data loginData
	buf := bytes.NewReader(ctx.PostBody())
	err := json.NewDecoder(buf).Decode(&data)
	if err != nil {
		errmsg("login Decode", err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if data.Username == "user" && data.Password == "userpass" {
		claims := jwtGo.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		}

		token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, claims)
		ss, err := token.SignedString(sKey)
		if err != nil {
			errmsg("login SignedString", err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.Set("Content-Type", "application/json")
		data := jsonData{
			Token: ss,
			Name:  data.Username,
			Admin: false,
		}
		err = json.NewEncoder(ctx).Encode(data)
		if err != nil {
			errmsg("login Encode", err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	} else {
		ctx.Error("Invalid Username or Password", fasthttp.StatusNotFound)
	}
}

func checkAuth(ctx *fasthttp.RequestCtx) bool {
	if cfg.Web.Auth {
		parser := &jwtGo.Parser{
			ValidMethods: []string{"HS256"},
		}
		bearer := parseBearerAuth(string(ctx.Request.Header.Peek("Authorization")))
		if bearer == "" {
			return false
		}
		token, err := parser.Parse(bearer, func(t *jwtGo.Token) (interface{}, error) {
			return sKey, nil
		})
		errChkMsg("checkAuth Parse", err)
		return token.Valid
	}
	return true
}
