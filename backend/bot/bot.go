package bot

import (
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/isiyar/daily-energy/backend/config"
)

func StartBot(c config.Config) {
	bot, err := tgbotapi.NewBotAPI(c.TelegramBotToken)
	if err != nil {
		log.Panicf("Failed to create bot: %v", err)
	}

	bot.Debug = c.Debug
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		switch update.Message.Command() {
		case "start":
			handleStartCommand(bot, update.Message.Chat.ID)
		}
	}
}

func handleStartCommand(bot *tgbotapi.BotAPI, chatID int64) {
	inlineBtn := tgbotapi.NewInlineKeyboardButtonWebApp(
		"Открыть Mini App",
		tgbotapi.WebAppInfo{URL: "https://test-srvr.ru"},
	)
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(inlineBtn),
	)

	msg := tgbotapi.NewMessage(chatID, "Привет! Добро пожаловать в Daily Energy!")
	msg.ReplyMarkup = inlineKeyboard

	if _, err := bot.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}