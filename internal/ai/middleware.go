package ai

import (
	"github.com/ab-dauletkhan/tg-ai-bot/internal/ai/gemini"
)

func GenerateResponse(prompt string, choosenAI string) string {
	switch choosenAI {
	case "gemini":
		resp, err := gemini.GenerateResponse(prompt)
		if err != nil {
			return "Failed to generate response from Gemini: " + err.Error()
		}
		return *resp
	default:
		return "No AI choosen"
	}
}
