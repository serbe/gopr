package main

import (
	"github.com/serbe/adb"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		return
	}
	db = adb.InitDB(
		cfg.Base.Dbname,
		"127.0.0.1",
		cfg.Base.User,
		cfg.Base.Password,
	)
	initServer(":"+cfg.Web.Port, cfg.Web.Log, cfg.Web.Auth)
}
