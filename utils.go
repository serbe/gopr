package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/serbe/adb"
)

var (
	logErrors bool
	db        *adb.ADB
)

// Config all vars
type Config struct {
	Web struct {
		Auth bool   `json:"auth"`
		Log  bool   `json:"log"`
		Port string `json:"port"`
	} `json:"web"`
	Base struct {
		// LogSQL   bool   `json:"logsql"`
		LogErr   bool   `json:"logerr"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Host     string `json:"host"`
	} `json:"base"`
}

func getConfig() (Config, error) {
	var c Config
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		errmsg("getConfig ReadFile", err)
		return c, err
	}
	if err = json.Unmarshal(file, &c); err != nil {
		errmsg("getConfig Unmarshal", err)
		return c, err
	}
	logErrors = c.Base.LogErr
	if c.Base.Name == "" {
		err = errors.New("Error: empty database name in config")
		errmsg("getConfig", err)
		return c, err
	}
	return c, err
}

func toInt(num string) int64 {
	id, err := strconv.ParseInt(num, 10, 64)
	errChkMsg("toInt", err)
	return id
}

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
