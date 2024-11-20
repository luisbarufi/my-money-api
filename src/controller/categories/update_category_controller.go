package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"github.com/luisbarufi/my-money-api/src/controller/categories/model/request"
	model "github.com/luisbarufi/my-money-api/src/model/categories"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) UpdateCategoryController(c *gin.Context) {
	logger.Info(
		"Init UpdateCategoryController",
		zap.String("journey", "updateCategory"),
	)

	categoryId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate category id, must be integer",
			err,
			zap.String("journey", "updateCategory"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	var categoryRequest request.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		logger.Error(
			"Error trying to validate category info",
			err,
			zap.String("journey", "updateCategory"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUpdateCategoryDomain(categoryRequest.CategoryName)

	if err := cc.categoryService.UpdateCategoryService(categoryId, domain); err != nil {
		logger.Error("Error calling UpdateCategoryService",
			err,
			zap.String("journey", "updateCategory"),
		)

		c.JSON(err.Code, err)

		return
	}

	logger.Info(
		"UpdateAccountController executed successfully",
		zap.String("categoryId", fmt.Sprintf("%d", categoryId)),
		zap.String("journey", "updateCategory"),
	)

	c.Status(http.StatusOK)
}
