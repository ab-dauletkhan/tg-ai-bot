package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateResponse(userRequest string) (*string, error) {
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
