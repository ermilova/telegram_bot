package main

import (
	"github.com/ermilova/telegram_bot/internal/service/product"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)
	productService := product.NewService()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/help - help\n"+
		"/list - list products")
	bot.Send(msg)
}
func listCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, service *product.Service) {
	outText := "Here all the products: \n\n"
	products := service.List()
	for _, p := range products {
		outText += p.Title
		outText += "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, outText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "You wrote "+message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
