package main

import (
	"github.com/atctx/get-mails-to-tgbot/internal/app"
	"github.com/atctx/get-mails-to-tgbot/internal/config"
)

func main() {
	app.NewApp(config.Cfg)
}
