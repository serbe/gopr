package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func startBot(token string) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/work", func(m *tb.Message) {
		var (
			arg int
		)
		text := strings.Trim(m.Text, " ")
		text = strings.Replace(text, "  ", " ", -1)
		split := strings.Split(text, " ")
		if len(split) == 2 {
			arg, _ = strconv.Atoi(split[1])
		}
		if arg > 100 {
			arg = 100
		} else if arg < 1 {
			arg = 1
		}
		list, err := db.ProxyGetRandomWorking(arg)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		b.Send(m.Sender, strings.Join(list, "\n"))
	})

	b.Handle("/anon", func(m *tb.Message) {
		var (
			arg int
		)
		text := strings.Trim(m.Text, " ")
		text = strings.Replace(text, "  ", " ", -1)
		split := strings.Split(text, " ")
		if len(split) == 2 {
			arg, _ = strconv.Atoi(split[1])
		}
		if arg > 100 {
			arg = 100
		} else if arg < 1 {
			arg = 1
		}
		list, err := db.ProxyGetRandomAnonymous(arg)
		if err != nil {
			b.Send(m.Sender, err.Error())
			return
		}
		b.Send(m.Sender, strings.Join(list, "\n"))
	})

	b.Start()
}
