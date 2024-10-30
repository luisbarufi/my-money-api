package request

type UserForgotPassword struct {
	Email string `json:"email" binding:"required,email"`
}

type UserUpdatePassword struct {
	Password string `json:"password" binding:"required,min=5,containsany=!@#$%&*"`
	Token    string `json:"token" binding:"required"`
}
