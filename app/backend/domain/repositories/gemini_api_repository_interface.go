package repositories

type IGeminiApiRepository interface {
	GetGeminiResponse(prompt string) (string, error)
}
