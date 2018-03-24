package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/go-chi/jwtauth"
// 	"github.com/go-chi/render"
// 	"github.com/valyala/fasthttp"
// )

// var sKey = []byte("aH0tH3P5up3RdYP3r53crEt")

// var tokenAuth *jwtauth.JWTAuth

// type loginData struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type jToken struct {
// 	Token string `json:"token"`
// 	Name  string `json:"name"`
// 	Admin bool   `json:"admin"`
// }

// func (l *loginData) Bind(r *http.Request) error {
// 	// just a post-process after a decode..
// 	return nil
// }

// func initAuth() {
// 	tokenString := "sdfksdkjfh87yKJFDuysdfuhkn.,sdesdfe2123568gasduyg6thjbkjJBF&*TgjhBFDUTygI&EDYhbekr*sdbf"

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return sKey, nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		fmt.Println(claims["foo"], claims["nbf"])
// 	} else {
// 		fmt.Println(err)
// 	}
// }

// func login(ctx *fasthttp.RequestCtx) {
// 	var data loginData
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		errmsg("login Decode", err)
// 		render.Status(r, http.StatusInternalServerError)
// 		render.PlainText(w, r, err.Error())
// 		err = r.Body.Close()
// 		errChkMsg("login Body.Close", err)
// 		return
// 	}

// 	if data.Username == "user" && data.Password == "userpass" {
// 		var tokenString string
// 		_, tokenString, err = tokenAuth.Encode(jwtauth.Claims{
// 			"admin": false,
// 			"name":  data.Username,
// 			"exp":   time.Now().Add(time.Hour * 24).Unix(),
// 		})
// 		if err != nil {
// 			errmsg("login SignedString", err)
// 			render.Status(r, http.StatusInternalServerError)
// 			render.PlainText(w, r, err.Error())
// 			err = r.Body.Close()
// 			errChkMsg("login Body.Close", err)
// 			return
// 		}
// 		render.JSON(w, r, jToken{
// 			Token: tokenString,
// 			Name:  data.Username,
// 			Admin: false,
// 		})
// 	} else {
// 		render.Status(r, http.StatusNotFound)
// 		render.PlainText(w, r, "Invalid Username or Password")
// 	}
// 	err = r.Body.Close()
// 	errChkMsg("login Body.Close", err)
// }

// // func login(w http.ResponseWriter, r *http.Request) {
// // 	var data loginData
// // 	err := json.NewDecoder(r.Body).Decode(&data)
// // 	if err != nil {
// // 		errmsg("login Decode", err)
// // 		render.Status(r, http.StatusInternalServerError)
// // 		render.PlainText(w, r, err.Error())
// // 		err = r.Body.Close()
// // 		errChkMsg("login Body.Close", err)
// // 		return
// // 	}

// // 	if data.Username == "user" && data.Password == "userpass" {
// // 		var tokenString string
// // 		_, tokenString, err = tokenAuth.Encode(jwtauth.Claims{
// // 			"admin": false,
// // 			"name":  data.Username,
// // 			"exp":   time.Now().Add(time.Hour * 24).Unix(),
// // 		})
// // 		if err != nil {
// // 			errmsg("login SignedString", err)
// // 			render.Status(r, http.StatusInternalServerError)
// // 			render.PlainText(w, r, err.Error())
// // 			err = r.Body.Close()
// // 			errChkMsg("login Body.Close", err)
// // 			return
// // 		}
// // 		render.JSON(w, r, jToken{
// // 			Token: tokenString,
// // 			Name:  data.Username,
// // 			Admin: false,
// // 		})
// // 	} else {
// // 		render.Status(r, http.StatusNotFound)
// // 		render.PlainText(w, r, "Invalid Username or Password")
// // 	}
// // 	err = r.Body.Close()
// // 	errChkMsg("login Body.Close", err)
// // }
