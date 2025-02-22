package dto

type GetBookRequestParam struct {
	Title  string `form:"title"`
	MaxNum int    `form:"maxNum,default=5"`
}
