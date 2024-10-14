package tgbot

import (
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func NewTgBot(token string) {
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("it's empty")
	})

	fmt.Println("tgbot started")
	b.Start()

}
