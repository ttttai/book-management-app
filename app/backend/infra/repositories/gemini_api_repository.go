package repositories

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/ttttai/golang/domain/repositories"
	"google.golang.org/api/option"
)

type GeminiApiRepository struct {
}

func NewGeminiApiRepository() repositories.IGeminiApiRepository {
	return &GeminiApiRepository{}
}

func (gr *GeminiApiRepository) GetGeminiResponse(prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	return convertResponse(resp), nil
}

func convertResponse(resp *genai.GenerateContentResponse) string {
	var builder strings.Builder

	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			builder.WriteString(fmt.Sprintf("%v", part))
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
