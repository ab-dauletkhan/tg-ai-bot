package main

import (
	"fmt"
	"log"
	"os"

	tgbot "github.com/ab-dauletkhan/tg-ai-bot/internals/tg-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Fatalf("Error creating new Bot API: %v", err)
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want implemented types of updates, so we'll ignore
		// any other types.
		if update.Message == nil && update.EditedMessage == nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I'm sorry, I can only respond to messages and edited messages.")
			msg.ReplyToMessageID = update.Message.MessageID

			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
			continue
		}

		if update.Message != nil {
			tgbot.PrintMessage(update.Message)

			if update.Message.MediaGroupID != "" {
				fmt.Println("not handled")

				// In case of MediaGroups we need to save them to the database and make a new reply MediaGroup
				// So that we will not answer separately to each media in the group
			}

			if update.Message.Photo != nil {
				photo := update.Message.Photo[0]

				msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileID(photo.FileID))

				if update.Message.Caption != "" {
					msg.Caption = "You sent a photo with a caption: " + update.Message.Caption
				} else {
					msg.Caption = "You sent a photo without a caption."
				}
				msg.ReplyToMessageID = update.Message.MessageID

				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
			} else {
				// Now that we know we've gotten a new message, we can construct a
				// reply! We'll take the Chat ID and Text from the incoming message
				// and use it to create a new message.
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You sent: "+update.Message.Text)
				// We'll also say that this message is a reply to the previous message.
				// For any other specifications than Chat ID or Text, you'll need to
				// set fields on the `MessageConfig`.
				msg.ReplyToMessageID = update.Message.MessageID

				// Okay, we're sending our message off! We don't care about the message
				// we just sent, so we'll discard it.
				if _, err := bot.Send(msg); err != nil {
					// Note that panics are a bad way to handle errors. Telegram can
					// have service outages or network errors, you should retry sending
					// messages or more gracefully handle failures.
					panic(err)
				}
			}

		} else if update.EditedMessage != nil {
			tgbot.PrintMessage(update.EditedMessage)
			// For the moment, edited message will get new reply message
			// Editing the previous message requires additional memory
			// to store the previous messages and time to traverse through them.
			//
			// This is will be implemented in the future.
			fmt.Printf("Edited message: %s\n", update.EditedMessage.Text)
			msg := tgbotapi.NewMessage(update.EditedMessage.Chat.ID, update.EditedMessage.Text)
			msg.ReplyToMessageID = update.EditedMessage.MessageID

			if _, err := bot.Send(msg); err != nil {
				fmt.Println("Panicking!!!!!!!!")
				panic(err)
			}
		}
	}
}
