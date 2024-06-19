package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote "+message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
