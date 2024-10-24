package validation

import (
	"encoding/json"
	"errors"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/rest_err"
	"github.com/luisbarufi/my-money-api/src/controller/model/request"
	"go.uber.org/zap"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}

func ValidateUserID(c *gin.Context) (uint64, *rest_err.RestErr) {
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Error trying to validate user id, must be integer",
			err,
			zap.String("journey", "deleteUser"),
		)
		return 0, rest_err.NewBadRequestError("Invalid user id")
	}
	return userId, nil
}

func ValidateUserEmail(c *gin.Context) (string, *rest_err.RestErr) {
	userEmail := c.Param("email")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate email", err, zap.String("journey", "findUserByEmail"))
		return "", rest_err.NewBadRequestError("Email is not valid")
	}
	return userEmail, nil
}

func ValidateUserUpdateRequest(c *gin.Context) (*request.UserUpdateRequest, *rest_err.RestErr) {
	var userRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "updateUser"))
		restErr := ValidateUserError(err)
		return nil, restErr
	}
	return &userRequest, nil
}

func ValidateUserInput(c *gin.Context) (
	*request.UserRequest, *rest_err.RestErr,
) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "createUser"),
		)
		restErr := ValidateUserError(err)
		return nil, restErr
	}
	return &userRequest, nil
}

func ValidateLoginUserInput(c *gin.Context) (
	*request.UserLogin, *rest_err.RestErr,
) {
	var userLogin request.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error(
			"Error trying to validate userLogin info",
			err,
			zap.String("journey", "createUser"),
		)
		restErr := ValidateUserError(err)
		return nil, restErr
	}
	return &userLogin, nil
}
