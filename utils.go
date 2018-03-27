package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/serbe/adb"
)

var (
	logErrors          bool
	db                 *adb.ADB
	cfg                Config
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

// Config all vars
type Config struct {
	Web struct {
		Auth bool `json:"auth"`
		// Log      bool   `json:"log"`
		CORS    bool   `json:"cors"`
		CorsURL string `json:"cors_url"`
		Port    string `json:"port"`
	} `json:"web"`
	Base struct {
		// LogSQL   bool   `json:"logsql"`
		LogErr   bool   `json:"logerr"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Host     string `json:"host"`
	} `json:"base"`
	Bot struct {
		Enable bool   `json:"enable"`
		Token  string `json:"token"`
	} `json:"bot"`
}

func getConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("getConfig ReadFile", err)
	}
	if err = json.Unmarshal(file, &cfg); err != nil {
		log.Fatal("getConfig Unmarshal", err)
	}
	logErrors = cfg.Base.LogErr
	if cfg.Base.Name == "" {
		err = errors.New("empty database name in config")
		log.Fatal("getConfig", err)
	}
}

// func toInt(num string) int64 {
// 	id, err := strconv.ParseInt(num, 10, 64)
// 	errChkMsg("toInt", err)
// 	return id
// }

func errmsg(str string, err error) {
	if logErrors {
		log.Println("Error in", str, err)
	}
}

func errChkMsg(str string, err error) {
	if logErrors && err != nil {
		log.Println("Error in", str, err)
	}
}

func getArgInt(str string) int {
	var result int
	text := strings.Trim(str, " ")
	text = strings.Replace(text, "  ", " ", -1)
	split := strings.Split(text, " ")
	if len(split) == 2 {
		result, _ = strconv.Atoi(split[1])
	}
	if result > 100 {
		result = 100
	} else if result < 1 {
		result = 1
	}
	return result
}

func getArgString(str string) string {
	var result string
	text := strings.Trim(str, " ")
	text = strings.Replace(text, "  ", " ", -1)
	split := strings.Split(text, " ")
	if len(split) == 2 {
		result = split[1]
	}
	return result
}

func parseBearerAuth(auth string) string {
	if strings.HasPrefix(auth, "Bearer ") {
		if bearer, err := base64.StdEncoding.DecodeString(auth[7:]); err == nil {
			return string(bearer)
		}
	}
	return ""
}
