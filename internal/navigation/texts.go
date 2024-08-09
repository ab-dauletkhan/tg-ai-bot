package navigation

import "fmt"

const (
	MainMenuText = `🤖 Welcome to GenAI Bot!

I can answer your questions using different AI models. You can choose your preferred AI and manage your balance for accessing premium features.

Available Commands:
- /help: Comprehensive guide
- /settings: Customize your experience

What would you like to do?`

	HelpText = `📖 Help Guide

1. Choose AI:
   - Select an AI model to answer your questions.
   - Some AIs are free, others require a balance.

2. Check Balance/Deposit:
   - View your current balance.
   - Deposit funds to use premium AI models.

3. Stop Conversation:
   - Use /stop to end the current conversation and return to the main menu.

4. Settings:
   - Customize your bot experience.

For more assistance, contact support.`

	SettingsText = `⚙️ Settings

Customize your GenAI Bot experience. More options coming soon!`

	StopText = `🛑 Conversation stopped. What would you like to do next?`

	DepositText = `💳 Deposit Balance

Please enter the amount you'd like to deposit:`
)

// AIModelText returns a message indicating the chosen AI model is now active.
func AIModelText(model string) string {
	return fmt.Sprintf(`✍️ AI Model %s is now active. Write your question below.

To stop the conversation, type /stop.`, model)
}

// BalanceText returns a message displaying the user's current balance.
func BalanceText(balance float64) string {
	return fmt.Sprintf(`💰 Your current balance is: %.2f.

What would you like to do next?`, balance)
}

// DepositSuccessText returns a message confirming a successful deposit and the new balance.
func DepositSuccessText(amount float64, newBalance float64) string {
	return fmt.Sprintf(`✅ Deposit successful! Your new balance is: %.2f.

What would you like to do next?`, newBalance)
}
