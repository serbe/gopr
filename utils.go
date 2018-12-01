package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	logErrors bool
	cfg       Config
)

// Config all vars
type Config struct {
	Log  bool   `json:"log"`
	Port string `json:"port"`
}

func getConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("getConfig ReadFile", err)
	}
	if err = json.Unmarshal(file, &cfg); err != nil {
		log.Fatal("getConfig Unmarshal", err)
	}
	logErrors = cfg.Log
}

func errChkMsg(str string, err error) {
	if logErrors && err != nil {
		log.Println("Error in", str, err)
	}
}
