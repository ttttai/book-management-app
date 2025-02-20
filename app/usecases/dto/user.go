package dto

type GetUserRequestParam struct {
	ID string `uri:"id" binding:"required"`
}

type CreateUserRequestParam struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
