package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ab-dauletkhan/tg-ai-bot/config"
	"github.com/ab-dauletkhan/tg-ai-bot/internal/telegram"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	cfg := config.LoadConfig()

	dbpool, err := telegram.InitDB(cfg.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	bot, err := telegram.NewBotAPI(cfg.TGToken)
	if err != nil {
		log.Fatalf("Error creating new Bot API: %v", err)
	}

	telegram.StartBot(bot, dbpool)
}
