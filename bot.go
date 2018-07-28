package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

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
		log.Println(arg, m.Text)
		var result string
		switch arg {
		case "":
			result = strconv.FormatInt(db.ProxyGetAllCount(), 10)
		case "work":
			result = strconv.FormatInt(db.ProxyGetAllWorkCount(), 10)
		case "anon":
			result = strconv.FormatInt(db.ProxyGetAllAnonymousCount(), 10)
		case "http":
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount("http"), 10)
		case "https":
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount("https"), 10)
		case "socks":
			result = strconv.FormatInt(db.ProxyGetAllSchemeCount("socks5"), 10)
		case "old":
			result = strconv.FormatInt(db.ProxyGetAllOldCount(), 10)
		default:
			result = "Use work, anon or empty string"
		}
		b.Send(m.Sender, result)
	})

	b.Handle("/countwork", func(m *tb.Message) {
		result := strconv.FormatInt(db.ProxyGetAllWorkCount(), 10)
		b.Send(m.Sender, result)
	})

	b.Handle("/countanon", func(m *tb.Message) {
		result := strconv.FormatInt(db.ProxyGetAllAnonymousCount(), 10)
		b.Send(m.Sender, result)
	})

	b.Start()
}
