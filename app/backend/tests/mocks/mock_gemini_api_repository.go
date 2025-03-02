package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockGeminiApiRepository struct {
	mock.Mock
}

func NewMockGeminiApiRepository() *MockGeminiApiRepository {
	return &MockGeminiApiRepository{}
}

func (m *MockGeminiApiRepository) GetGeminiResponse(prompt string) (string, error) {
	args := m.Called(prompt)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}
