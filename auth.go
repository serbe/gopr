package main

import (
	"encoding/json"
	"net/http"

	jwtGo "github.com/dgrijalva/jwt-go"
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

func login(w http.ResponseWriter, req *http.Request) {
	var data loginData
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		errmsg("login Decode", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		data := jsonData{
			Token: ss,
			Name:  data.Username,
			Admin: false,
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			errmsg("login Encode", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid Username or Password", http.StatusNotFound)
	}
}

func checkAuth(w http.ResponseWriter, req *http.Request) bool {
	if cfg.Web.Auth {
		parser := &jwtGo.Parser{
			ValidMethods: []string{"HS256"},
		}
		bearer := parseBearerAuth(string(req.Header.Get("Authorization")))
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
