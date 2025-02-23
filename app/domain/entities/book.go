package entities

type Book struct {
	ID                int
	ISBN              int
	TitleName         string
	TitleNameKana     string
	PublisherName     string
	PublisherNameKana string
	PublishDate       *string
	Price             int
}
