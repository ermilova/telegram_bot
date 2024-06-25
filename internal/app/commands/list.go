package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	outText := "Here all the products: \n\n"
	products := c.service.List()
	for _, p := range products {
		outText += p.Title
		outText += "\n"
	}
	serializedData, _ := json.Marshal(CommandData{
		Offset: 10,
	})
	var keyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Nexp page", string(serializedData)),
		),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, outText)
	msg.ReplyMarkup = keyboard
	c.bot.Send(msg)
}
func init() {
	registeredCommands["list"] = (*Commander).List
}
