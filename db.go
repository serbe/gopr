package main

import (
	"log"
	"time"

	"github.com/go-pg/pg"
)

var dbase *pg.DB

// Proxy - proxy unit
type Proxy struct {
	// Insert   bool          `sql:"-"           json:"-"`
	// Update   bool          `sql:"-"           json:"-"`
	// URL      *url.URL      `sql:"-"           json:"-"`
	ID        int64         `sql:"id,pk"       json:"id"`
	Hostname  string        `sql:"hostname"    json:"-"`
	Host      string        `sql:"host"        json:"host"`
	Port      string        `sql:"port"        json:"port"`
	IsWork    bool          `sql:"work"        json:"work"`
	IsAnon    bool          `sql:"anon"        json:"anon"`
	Checks    int           `sql:"checks"      json:"check"`
	CreateStr string        `sql:"-"           json:"create"`
	UpdateSrt string        `sql:"-"           json:"update"`
	CreateAt  time.Time     `sql:"create_at"   json:"-"`
	UpdateAt  time.Time     `sql:"update_at"   json:"-"`
	Response  time.Duration `sql:"response"    json:"response"`
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
func initDB(
	dbname string,
	user string,
	password string,
	logsql bool,
) {
	opt := pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	}
	dbase = pg.Connect(&opt)
	if logsql {
		dbase.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}
	_, err := dbase.Exec("SELECT")
	if err != nil {
		panic(err)
	}
}

func getProxyByID(id int64) Proxy {
	var p Proxy
	err := dbase.Model(&p).Where("id = ?", id).Select()
	errchkmsg("getProxy", err)
	return p
}

func getProxies(p string) ([]Proxy, error) {
	var (
		ps  []Proxy
		err error
	)
	if p == "" {
		err = dbase.Model(&ps).Select()
	} else {
		err = dbase.Model(&ps).Where(p).Select()
	}
	if err != nil {
		return ps, err
	}
	for i, p := range ps {
		ps[i].CreateStr = p.CreateAt.Format("02.01.2006")
		ps[i].UpdateSrt = p.UpdateAt.Format("02.01.2006")
	}
	return ps, nil
}

func getAllProxies() []Proxy {
	ps, err := getProxies("")
	errchkmsg("getAllProxies", err)
	return ps
}

func getAllWorkProxies() []Proxy {
	ps, err := getProxies("work = TRUE")
	errchkmsg("getAllWorkProxies", err)
	return ps
}

func getAllAnonProxies() []Proxy {
	ps, err := getProxies("work = TRUE AND anon = TRUE")
	errchkmsg("getAllAnonProxies", err)
	return ps
}

func getAllCount() int64 {
	var ps []Proxy
	c, err := dbase.Model(&ps).Count()
	errchkmsg("getAllProxies", err)
	return int64(c)
}

func getAllWorkCount() int64 {
	var ps []Proxy
	c, err := dbase.Model(&ps).Where("work = TRUE").Count()
	errchkmsg("getAllWorkProxies", err)
	return int64(c)
}

func getAllAnonCount() int64 {
	var ps []Proxy
	c, err := dbase.Model(&ps).Where("work = TRUE AND anon = TRUE").Count()
	errchkmsg("getAllAnonProxies", err)
	return int64(c)
}
