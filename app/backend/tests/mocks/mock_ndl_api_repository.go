package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/ttttai/golang/domain/entities"
)

type MockNdlApiRepository struct {
	mock.Mock
}

func NewMockNdlApiRepository() *MockNdlApiRepository {
	return &MockNdlApiRepository{}
}

func (m *MockNdlApiRepository) GetBooksFromNdlApi(title string, maxNum int, offSet int) (*[]entities.BookInfo, error) {
	args := m.Called(title, maxNum)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]entities.BookInfo), args.Error(1)
}
