package dto

type GetByEmailRequestParam struct {
	Email string `form:"email"`
}

type CreateAuthenticationRequestParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
