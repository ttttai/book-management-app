package dto

type GetBookRequestParam struct {
	Title  string `form:"title"`
	MaxNum int    `form:"maxNum,default=5"`
}

type GetBookInfoRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

type DeleteBookRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

type CreateBookInfoRequestParam struct {
	Book     BookParam      `json:"book"`
	Authors  []AuthorParam  `json:"authors"`
	Subjects []SubjectParam `json:"subjects"`
}

type BookParam struct {
	ISBN              int    `json:"isbn" binding:"required"`
	TitleName         string `json:"title_name"`
	TitleNameKana     string `json:"title_name_kana"`
	PublisherName     string `json:"publisher_name"`
	PublisherNameKana string `json:"publisher_name_kana"`
	PublishDate       string `json:"publish_date"`
	Price             int    `json:"price"`
	Status            int    `json:"status"`
}

type AuthorParam struct {
	Name     string `json:"name"`
	NameKana string `json:"name_kana"`
}

type SubjectParam struct {
	SubjectName string `json:"subject_name"`
	SubjectKana string `json:"subject_kana"`
}

type UpdateBookRequestParam struct {
	ID                int    `json:"id" binding:"required"`
	ISBN              int    `json:"isbn" binding:"required"`
	TitleName         string `json:"title_name"`
	TitleNameKana     string `json:"title_name_kana"`
	PublisherName     string `json:"publisher_name"`
	PublisherNameKana string `json:"publisher_name_kana"`
	PublishDate       string `json:"publish_date"`
	Price             int    `json:"price"`
}

type UpdateBookStatusRequestPathParam struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateBookStatusRequestBodyParam struct {
	Status int `json:"status"`
}
