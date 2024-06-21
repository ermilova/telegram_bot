package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Print("Wrong args", args)
		return
	}
	product, err := c.service.Get(idx)
	if err != nil {
		log.Print("Failed to get product", idx)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, product.Title)
	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
func init() {
	registeredCommands["get"] = (*Commander).Get
}
