package app

import (
	"github.com/atctx/get-mails-to-tgbot/internal/config"
	"github.com/atctx/get-mails-to-tgbot/internal/tgbot"
)

func NewApp(cfg config.Config) {
	tgbot.NewTgBot(cfg.TgToken)
}
