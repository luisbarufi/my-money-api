package response

type PasswordForgotResponse struct {
	User       interface{} `json:"user"`
	ResetToken string      `json:"resetToken"`
}
