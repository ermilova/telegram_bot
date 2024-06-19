package commands

import (
	"github.com/ermilova/telegram_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, service *product.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: service,
	}
}
