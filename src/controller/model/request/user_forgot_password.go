package request

type UserForgotPassword struct {
	Email string `json:"email" binding:"required,email"`
}
