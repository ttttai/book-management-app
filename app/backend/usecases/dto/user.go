package dto

type GetUserRequestParam struct {
	ID string `uri:"id" binding:"required"`
}

type CreateUserRequestParam struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserRequestPathParam struct {
	ID string `uri:"id" binding:"required"`
}

type UpdateUserRequestBodyParam struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type DeleteUserRequestParam struct {
	ID string `uri:"id" binding:"required"`
}

type GetUserByNameRequestParam struct {
	Name string `form:"name" binding:"required"`
}
