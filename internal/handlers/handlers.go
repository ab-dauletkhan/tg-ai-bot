package handlers

import (
	"fmt"
	"log"

	"github.com/ab-dauletkhan/tg-ai-bot/internal/balance"
	"github.com/ab-dauletkhan/tg-ai-bot/internal/keyboards"
	"github.com/ab-dauletkhan/tg-ai-bot/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, dbpool *pgxpool.Pool) {
	if message.IsCommand() {
		utils.PrintMessage(message)

		switch message.Command() {
		// case "help":
		// 	msg.Text = "Help command received."
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, HandleStart(bot, message, dbpool))
			msg.ReplyMarkup = keyboards.MainMenuKeyboard()
			msg.ParseMode = tgbotapi.ModeMarkdown
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}

			// case "balance":
			// 	HandleBalance(bot, message, dbpool, msg)
			// default:
			// 	msg.Text = "I don't know that command"
		}
	} else {
		utils.PrintMessage(message)
		HandleNonCommandMessage(bot, message)
	}
}

func HandleEditedMessage(bot *tgbotapi.BotAPI, editedMessage *tgbotapi.Message) {
	utils.PrintMessage(editedMessage)
	fmt.Printf("Edited message: %s\n", editedMessage.Text)
	msg := tgbotapi.NewMessage(editedMessage.Chat.ID, editedMessage.Text)
	msg.ReplyToMessageID = editedMessage.MessageID

	if _, err := bot.Send(msg); err != nil {
		fmt.Println("Panicking!!!!!!!!")
		panic(err)
	}
}

func HandleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message, dbpool *pgxpool.Pool) string {
	welcomeMessage :=
		`# 🤖 Welcome to GenAI Bot!

		_I can_ answer your questions using different AI models. You can choose your preferred AI and manage your balance for accessing premium features.

		Available Commands:
		- /help: Comprehensive guide
		- /settings: Customize your experience

		What would you like to do?
		`

	userName := message.From.FirstName + " " + message.From.LastName
	resp := fmt.Sprintf("Hello, %s! I'm here to assist you.\n\n", userName)

	b, err := balance.GetBalanceByUserID(dbpool, message.From.ID)
	if err != nil && err != pgx.ErrNoRows {
		resp += "Error checking balance: " + err.Error()
	} else if b == nil {
		b = &balance.Balance{
			UserID: message.From.ID,
			Amount: 0.0,
			Status: "active",
		}

		err := balance.InsertBalance(dbpool, b)
		if err != nil {
			resp += "Error inserting balance: " + err.Error()
		} else {
			fmt.Println("Balance inserted")
			resp += "Your balance has been set to 0.0."
		}
	} else {
		fmt.Println("Balance already exists")
		resp += fmt.Sprintf("Your current balance is %.2f.", b.Amount)
	}

	return welcomeMessage + resp
}

func HandleBalance(bot *tgbotapi.BotAPI, message *tgbotapi.Message, dbpool *pgxpool.Pool, msg tgbotapi.MessageConfig) {
	b, err := balance.GetBalanceByUserID(dbpool, message.From.ID)
	if err != nil {
		msg.Text = "Error getting balance: " + err.Error()
	} else {
		msg.Text = fmt.Sprintf("Your balance is %.2f", b.Amount)
	}
}

func HandleNonCommandMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if message.Photo != nil {
		HandlePhotoMessage(bot, message)
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "You sent: "+message.Text)
		msg.ReplyToMessageID = message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}

func HandlePhotoMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	photo := message.Photo[0]
	msg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileID(photo.FileID))

	if message.Caption != "" {
		msg.Caption = "You sent a photo with a caption: " + message.Caption
	} else {
		msg.Caption = "You sent a photo without a caption."
	}
	msg.ReplyToMessageID = message.MessageID

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

// HandleCallbackQuery processes the callback queries triggered by inline keyboard buttons.
func HandleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery, dbpool *pgxpool.Pool) {
	var msg tgbotapi.EditMessageTextConfig

	switch callbackQuery.Data {
	case "choose_ai":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			callbackQuery.Message.Chat.ID,
			callbackQuery.Message.MessageID,
			"Please choose an AI model.",
			keyboards.ChooseAIKeyboard(),
		)

	case "check_balance":
		b, err := balance.GetBalanceByUserID(dbpool, callbackQuery.From.ID)
		if err != nil {
			msg = tgbotapi.NewEditMessageText(
				callbackQuery.Message.Chat.ID,
				callbackQuery.Message.MessageID,
				"Error getting balance: "+err.Error(),
			)
		} else {
			msg = tgbotapi.NewEditMessageText(
				callbackQuery.Message.Chat.ID,
				callbackQuery.Message.MessageID,
				fmt.Sprintf("Your current balance is %.2f.", b.Amount),
			)
		}

	case "back_to_main_menu":
		msg = tgbotapi.NewEditMessageTextAndMarkup(
			callbackQuery.Message.Chat.ID,
			callbackQuery.Message.MessageID,
			"Back to the main menu.",
			keyboards.MainMenuKeyboard(),
		)

	default:
		msg = tgbotapi.NewEditMessageText(
			callbackQuery.Message.Chat.ID,
			callbackQuery.Message.MessageID,
			"Unknown command.",
		)
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
