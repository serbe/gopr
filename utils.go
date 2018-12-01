package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/serbe/adb"
)

var (
	logErrors bool
	db        *adb.ADB
	cfg       Config
)

// Config all vars
type Config struct {
	Web struct {
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

func errChkMsg(str string, err error) {
	if logErrors && err != nil {
		log.Println("Error in", str, err)
	}
}
