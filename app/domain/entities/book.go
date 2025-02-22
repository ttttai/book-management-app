package entities

type Book struct {
	ID                uint
	ISBN              int
	TitleName         string
	TitleNameKana     string
	PublisherName     string
	PublisherNameKana string
	PublishDate       *string
	Price             int
}
