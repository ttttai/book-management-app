package services

import (
	"strings"

	"github.com/ttttai/golang/domain/repositories"
)

type IGeminiApiService interface {
	GetGeminiResponse(prompt string) ([]string, error)
}

type GeminiApiService struct {
	geminiApiRepository repositories.IGeminiApiRepository
}

func NewGeminiApiService(geminiApiRepository repositories.IGeminiApiRepository) IGeminiApiService {
	return &GeminiApiService{
		geminiApiRepository: geminiApiRepository,
	}
}

func (gs *GeminiApiService) GetGeminiResponse(prompt string) ([]string, error) {
	result, err := gs.geminiApiRepository.GetGeminiResponse(prompt)
	if err != nil {
		return []string{}, err
	}

	return getBookTitles(result), nil
}

func getBookTitles(res string) []string {
	var bookTitles []string

	lines := strings.Split(res, "\n")
	for _, line := range lines {
		title := strings.TrimSpace(line)
		title = strings.TrimPrefix(title, "*")
		title = strings.Trim(title, "* ")

		if title != "" {
			bookTitles = append(bookTitles, title)
		}
	}

	// 1文目は「読書履歴に基づき、おすすめの本を5冊提案します。」のため除外
	return bookTitles[1:]
}
