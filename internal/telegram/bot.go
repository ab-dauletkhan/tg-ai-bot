package telegram

import (
	"context"
	"log"

	"github.com/ab-dauletkhan/tg-ai-bot/internal/handlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(dbURL string) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dbURL)
}

func NewBotAPI(token string) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot, nil
}

func StartBot(bot *tgbotapi.BotAPI, dbpool *pgxpool.Pool) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			handlers.HandleMessage(bot, update.Message, dbpool)
		} else if update.EditedMessage != nil {
			handlers.HandleEditedMessage(bot, update.EditedMessage)
		} else if update.CallbackQuery != nil {
			handlers.HandleCallbackQuery(bot, update.CallbackQuery, dbpool)
		}
	}
}
