package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const cAnon = "anon"
const cHTTP = "http"
const cHTTPS = "https"
const cSocks = "socks"
const cSocks5 = "socks5"
const cOld = "old"
const cWork = "work"

func startBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.Bot.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/work", func(m *tb.Message) {
		list, err := db.ProxyGetRandomWorking(getArgInt(m.Text))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		b.Send(m.Sender, strings.Join(list, "\n"))
	})

	b.Handle("/anon", func(m *tb.Message) {
		list, err := db.ProxyGetRandomWorking(getArgInt(m.Text))
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		b.Send(m.Sender, strings.Join(list, "\n"))
	})

	b.Handle("/count", func(m *tb.Message) {
		arg := getArgString(m.Text)
		var result string
		switch arg {
		case "":
			result = strconv.FormatInt(db.ProxyGetAllCount(), 10)
		case cWork:
			result = strconv.FormatInt(db.ProxyGetAllWorkCount(), 10)
		case cAnon:
			result = strconv.FormatInt(db.ProxyGetAllAnonymousCount(), 10)
		case cHTTP:
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount(cHTTP), 10)
		case cHTTPS:
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount(cHTTPS), 10)
		case cSocks:
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount(cSocks5), 10)
		case cOld:
			result = strconv.FormatInt(db.ProxyGetAllOldCount(), 10)
		default:
			result = "Use work, anon, http, https, socks5, old or empty string"
		}
		b.Send(m.Sender, result)
	})

	b.Handle("/countwork", func(m *tb.Message) {
		arg := getArgString(m.Text)
		var result string
		switch arg {
		case "":
			result = strconv.FormatInt(db.ProxyGetAllWorkCount(), 10)
		case cHTTP:
			result = strconv.FormatInt(db.ProxyGetAllWorkingSchemeCount(cHTTP), 10)
		case cHTTPS:
			result = strconv.FormatInt(db.ProxyGetAllWorkingSchemeCount(cHTTPS), 10)
		case cSocks:
			result = strconv.FormatInt(db.ProxyGetAllWorkingSchemeCount(cSocks5), 10)
		default:
			result = "Use http, https, socks5 or empty string"
		}
		b.Send(m.Sender, result)
	})

	b.Handle("/countanon", func(m *tb.Message) {
		arg := getArgString(m.Text)
		var result string
		switch arg {
		case "":
			result = strconv.FormatInt(db.ProxyGetAllAnonymousCount(), 10)
		case cHTTP:
			result = strconv.FormatInt(db.ProxyGetAllAnonymousSchemeCount(cHTTP), 10)
		case cHTTPS:
			result = strconv.FormatInt(db.ProxyGetAllAnonymousSchemeCount(cHTTPS), 10)
		case cSocks:
			result = strconv.FormatInt(db.ProxyGetAllAnonymousSchemeCount(cSocks5), 10)
		default:
			result = "Use http, https, socks5 or empty string"
		}
		b.Send(m.Sender, result)
	})

	b.Start()
}
