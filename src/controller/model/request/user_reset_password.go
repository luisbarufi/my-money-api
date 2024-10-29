package request

type UserResetPassword struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=5,containsany=!@#$%&*"`
}
