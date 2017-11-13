package main

import (
	"log"
	"net/url"
	"time"

	"github.com/go-pg/pg"
)

var DB *pg.DB

// Proxy - proxy unit
type Proxy struct {
	Insert   bool          `sql:"-"           json:"-"`
	Update   bool          `sql:"-"           json:"-"`
	Hostname string        `sql:"hostname,pk" json:"hostname"`
	URL      *url.URL      `sql:"-"           json:"-"`
	Host     string        `sql:"host"        json:"-"`
	Port     string        `sql:"port"        json:"-"`
	IsWork   bool          `sql:"work"        json:"-"`
	IsAnon   bool          `sql:"anon"        json:"-"`
	Checks   int           `sql:"checks"      json:"-"`
	CreateAt time.Time     `sql:"create_at"   json:"-"`
	UpdateAt time.Time     `sql:"update_at"   json:"-"`
	Response time.Duration `sql:"response"    json:"-"`
}

// // Link - link unit
// type Link struct {
// 	Insert   bool      `sql:"-"           json:"-"`
// 	Update   bool      `sql:"-"           json:"-"`
// 	Hostname string    `sql:"hostname,pk" json:"hostname"`
// 	UpdateAt time.Time `sql:"update_at"   json:"-"`
// 	Iterate  bool      `sql:"iterate"     json:"-"`
// 	Num      int64     `sql:"num"         json:"-"`
// }

// initDB initialize database
func initDB(host string, dbname string, user string, password string, logsql bool) {
	opt := pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	}
	if host != "" {
		opt.Addr = host
	}
	DB = pg.Connect(&opt)
	if logsql {
		DB.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}
}

func getProxyByID(id int64) Proxy {
	var p Proxy
	err := DB.Model(&p).Where("id = ?", id).Select()
	errchkmsg("getProxy", err)
	return p
}

func getAllProxies() []Proxy {
	var ps []Proxy
	err := DB.Model(&ps).Select()
	errchkmsg("getAllProxies", err)
	return ps
}

func getAllWorkProxies() []Proxy {
	var ps []Proxy
	err := DB.Model(&ps).Where("work = TRUE").Select()
	errchkmsg("getAllWorkProxies", err)
	return ps
}

func getAllAnonProxies() []Proxy {
	var ps []Proxy
	err := DB.Model(&ps).Where("work = TRUE AND anon = TRUE").Select()
	errchkmsg("getAllAnonProxies", err)
	return ps
}
