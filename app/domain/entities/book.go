package entities

type Book struct {
	ID                uint
	ISBN              int
	TitleName         string
	TitleKana         string
	Authors           []Author
	PublisherName     string
	PublisherNameKana string
	PublishDate       string
	Subjects          []Subject
	Price             int
}

type Author struct {
	Name     string
	NameKana string
}

type Subject struct {
	SubjectName string
	SubjectKana string
}
