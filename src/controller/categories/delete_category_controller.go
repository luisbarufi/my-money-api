package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/configuration/logger"
	"github.com/luisbarufi/my-money-api/src/configuration/validation"
	"go.uber.org/zap"
)

func (cc *categoryControllerInterface) DeleteCategoryController(c *gin.Context) {
	logger.Info(
		"Init DeleteCategoryController",
		zap.String("journey", "deleteCategory"),
	)

	categoryId, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		logger.Error(
			"Error trying to validate category id, must be integer",
			err,
			zap.String("journey", "deleteCategory"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	if err := cc.categoryService.DeleteCategoryService(categoryId); err != nil {
		logger.Error(
			"Error calling DeleteCategoryService",
			err,
			zap.String("journey", "deleteCategory"),
		)

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	logger.Info("DeleteCategoryController executed successfully",
		zap.String("categoryId", fmt.Sprintf("%d", categoryId)),
		zap.String("journey", "deleteCategory"),
	)

	c.Status(http.StatusOK)
}
