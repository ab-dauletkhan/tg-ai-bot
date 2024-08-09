package navigation

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	chooseAIButton := tgbotapi.NewInlineKeyboardButtonData("Choose AI", "choose_ai")
	checkBalanceButton := tgbotapi.NewInlineKeyboardButtonData("Check Balance/Deposit", "check_balance")

	row := tgbotapi.NewInlineKeyboardRow(chooseAIButton, checkBalanceButton)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	return keyboard
}

func ChooseAIKeyboard() tgbotapi.InlineKeyboardMarkup {
	geminiButton := tgbotapi.NewInlineKeyboardButtonData("Gemini", "gemini")
	backToMainMenuButton := tgbotapi.NewInlineKeyboardButtonData("Back to Main Menu", "back_to_main_menu")

	row1 := tgbotapi.NewInlineKeyboardRow(geminiButton, backToMainMenuButton)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1)

	return keyboard
}

func CheckBalanceKeyboard() tgbotapi.InlineKeyboardMarkup {
	depositButton := tgbotapi.NewInlineKeyboardButtonURL("Deposit", "https://google.com")
	backToMainMenuButton := tgbotapi.NewInlineKeyboardButtonData("Back to Main Menu", "back_to_main_menu")

	row := tgbotapi.NewInlineKeyboardRow(depositButton, backToMainMenuButton)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	return keyboard
}
