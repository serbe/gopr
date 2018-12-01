package main

import "github.com/serbe/adb"

func main() {
	getConfig()
	db = adb.InitDB(
		cfg.Base.Name,
		cfg.Base.Host,
		cfg.Base.User,
		cfg.Base.Password,
	)
	initServer()
}
