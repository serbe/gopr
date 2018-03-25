package main

import (
	"bytes"
	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
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
	err := json.NewDecoder(bytes.NewReader(ctx.PostBody())).Decode(&data)
	if err != nil {
		errmsg("login Decode", err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if data.Username == "user" && data.Password == "userpass" {
		claims := jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(sKey)
		if err != nil {
			errmsg("login SignedString", err)
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
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
		parser := &jwt.Parser{
			ValidMethods: []string{"HS256"},
		}
		bearer := parseBearerAuth(string(ctx.Request.Header.Peek("Authorization")))
		if bearer == "" {
			return false
		}
		token, err := parser.Parse(bearer, func(t *jwt.Token) (interface{}, error) {
			return sKey, nil
		})
		errChkMsg("checkAuth Parse", err)
		return token.Valid
	}
	return true
}
