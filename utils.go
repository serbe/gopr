package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
)

var (
	logErrors bool
)

// Config all vars
type Config struct {
	Web struct {
		Auth bool   `json:"auth"`
		Log  bool   `json:"log"`
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"web"`
	Base struct {
		LogSQL   bool   `json:"logsql"`
		LogErr   bool   `json:"logerr"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	} `json:"base"`
}

func getConfig() (c Config, err error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		errmsg("getConfig ReadFile", err)
		return
	}
	if err = json.Unmarshal(file, &c); err != nil {
		errmsg("getConfig Unmarshal", err)
		return
	}
	logErrors = c.Base.LogErr
	if c.Base.Dbname == "" {
		err := errors.New("Error: empty database name in config")
		errmsg("getConfig", err)
		return c, err
	}
	return
}

func toInt(num string) int64 {
	id, err := strconv.ParseInt(num, 10, 64)
	errchkmsg("toInt", err)
	return id
}

func errmsg(str string, err error) {
	if logErrors {
		log.Println("Error in", str, err)
	}
}

func errchkmsg(str string, err error) {
	if logErrors && err != nil {
		log.Println("Error in", str, err)
	}
}
