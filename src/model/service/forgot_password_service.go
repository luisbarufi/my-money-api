package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/mailer"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) ForgotPasswordService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init reset password services", zap.String("journey", "resetPasswordService"))

	user, err := ud.FindUserByEmailService(userDomain.GetEmail())
	if err != nil {
		logger.Error(
			"Error calling findUserByEmailAndPasswordServices",
			err,
			zap.String("journey", "resetPasswordService"),
		)
		return nil, "", err
	}

	resetToken, err := user.GenerateResetToken()
	if err != nil {
		logger.Error(
			"Error generating jwt token",
			err,
			zap.String("journey", "resetPasswordService"),
		)
		return nil, "", err
	}

	baseUrl := env.GetEnv("BASE_URL")
	forgotPasswordURL := fmt.Sprintf("%supdate-password?token=%s", baseUrl, resetToken)

	smtpPort, rest_err := strconv.Atoi(env.GetEnv("SMTP_PORT"))
	if rest_err != nil {
		log.Fatalf("Error trying to connect to SMTP port, error: %s\n", err.Error())
		return nil, "", nil
	}

	mailService := mailer.NewSMTPMailService(mailer.SMTPConfig{
		Host:     env.GetEnv("SMTP_HOST"),
		Port:     smtpPort,
		UserName: env.GetEnv("SMTP_USERNAME"),
		Password: env.GetEnv("SMTP_PASSWORD"),
		From:     env.GetEnv("SMTP_FROM"),
	})

	mailService.Send(mailer.MailMessage{
		To:      []string{"jhon.doe@example.com"},
		Subject: "Test email",
		Body:    []byte(forgotPasswordURL),
	})

	logger.Info(
		"ResetPasswordService services executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.GetID())),
		zap.String("journey", "resetPasswordService"),
	)
	return user, resetToken, nil
}
