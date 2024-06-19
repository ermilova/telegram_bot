package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	outText := "Here all the products: \n\n"
	products := c.service.List()
	for _, p := range products {
		outText += p.Title
		outText += "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, outText)
	c.bot.Send(msg)
}
