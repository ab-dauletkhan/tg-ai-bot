package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GeminiRandR(userRequest string) (*string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, fmt.Errorf("failed to create AI client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(userRequest))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %v", err)
	}

	str_resp, err := GetTextFromPart(resp.Candidates[0].Content)
	if err != nil {
		return nil, fmt.Errorf("failed to get text from part: %v", err)
	}

	return &str_resp, nil
}

// GetTextFromPart extracts the string response from a Part if it is of type Text
func GetTextFromPart(content *genai.Content) (string, error) {
	switch p := content.Parts[0].(type) {
	case genai.Text:
		return string(p), nil
	default:
		return "", fmt.Errorf("part is not of type Text")
	}
}

/*
package main

import (
	"log"
	"os"
	"time"

	"github.com/ab-dauletkhan/tg-ai-bot/internals/gemini"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	pref := tele.Settings{
		Token:  os.Getenv("TG_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! I will repeat after you.")
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Just type something and I will repeat after you.")
	})

	b.Handle("/settings", func(c tele.Context) error {
		return c.Send("Settings are not available yet.")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		resp, err := gemini.GeminiRandR(c.Text())
		if err != nil {
			return c.Send("Failed to generate content: " + err.Error())
		}
		return c.Send(*resp)
	})

	b.Start()
}

*/
