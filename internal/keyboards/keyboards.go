package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func MainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	chooseAIButton := tgbotapi.NewInlineKeyboardButtonData("Choose AI", "choose_ai")
	checkBalanceButton := tgbotapi.NewInlineKeyboardButtonData("Check Balance/Deposit", "check_balance")

	row := tgbotapi.NewInlineKeyboardRow(chooseAIButton, checkBalanceButton)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	return keyboard
}

func ChooseAIKeyboard() tgbotapi.InlineKeyboardMarkup {
	gpt3Button := tgbotapi.NewInlineKeyboardButtonData("GPT-3", "gpt3")
	gpt4Button := tgbotapi.NewInlineKeyboardButtonData("GPT-4", "gpt4")
	gpt5Button := tgbotapi.NewInlineKeyboardButtonData("GPT-5", "gpt5")
	backToMainMenuButton := tgbotapi.NewInlineKeyboardButtonData("Back to Main Menu", "back_to_main_menu")

	row1 := tgbotapi.NewInlineKeyboardRow(gpt3Button, gpt4Button)
	row2 := tgbotapi.NewInlineKeyboardRow(gpt5Button, backToMainMenuButton)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	return keyboard
}
