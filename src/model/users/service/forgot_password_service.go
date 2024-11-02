package service

import (
	"fmt"
	"strconv"

	"github.com/luisbarufi/my-money-api/src/configuration/env"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/mailer"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	model "github.com/luisbarufi/my-money-api/src/model/users"

	"go.uber.org/zap"
)

var (
	BASE_URL      = "BASE_URL"
	SMTP_HOST     = "SMTP_HOST"
	SMTP_USERNAME = "SMTP_USERNAME"
	SMTP_PASSWORD = "SMTP_PASSWORD"
	SMTP_FROM     = "SMTP_FROM"
)

func (ud *userDomainService) ForgotPasswordService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info(
		"Init ForgotPasswordService",
		zap.String("journey", "forgotPassword"),
	)

	user, err := ud.FindUserByEmailService(userDomain.GetEmail())

	if err != nil {
		logger.Error(
			"Error calling FindUserByEmailService",
			err,
			zap.String("journey", "forgotPassword"),
		)

		return nil, "", err
	}

	resetToken, err := user.GenerateResetToken()

	if err != nil {
		logger.Error(
			"Error generating jwt token",
			err,
			zap.String("journey", "forgotPassword"),
		)

		return nil, "", err
	}

	baseUrl := env.GetEnv(BASE_URL)

	forgotPasswordURL := fmt.Sprintf("%supdate-password?token=%s", baseUrl, resetToken)

	smtpPort, rest_err := strconv.Atoi(env.GetEnv("SMTP_PORT"))

	if rest_err != nil {
		logger.Error(
			"Error trying to connect to SMTP port",
			err,
			zap.String("journey", "forgotPassword"),
		)

		return nil, "", nil
	}

	mailService := mailer.NewSMTPMailService(mailer.SMTPConfig{
		Host:     env.GetEnv(SMTP_HOST),
		Port:     smtpPort,
		UserName: env.GetEnv(SMTP_USERNAME),
		Password: env.GetEnv(SMTP_PASSWORD),
		From:     env.GetEnv(SMTP_FROM),
	})

	mailService.Send(mailer.MailMessage{
		To:      []string{"jhon.doe@example.com"},
		Subject: "Test email",
		Body:    []byte(forgotPasswordURL),
	})

	logger.Info("Email sent successfully", zap.String("journey", "forgotPassword"))

	logger.Info(
		"ForgotPasswordService executed successfully",
		zap.String("userId", fmt.Sprintf("%d", user.GetID())),
		zap.String("journey", "forgotPassword"),
	)

	return user, resetToken, nil
}
