package main

import (
	"log"

	"github.com/serbe/adb"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Panic(err)
	}
	db = adb.InitDB(
		cfg.Base.Name,
		cfg.Base.Host,
		cfg.Base.User,
		cfg.Base.Password,
	)
	initServer(":"+cfg.Web.Port, cfg.Web.Log, cfg.Web.Auth)
}
