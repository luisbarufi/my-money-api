package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luisbarufi/my-money-api/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/users/:id", controller.FindUserByID)
	r.GET("/users/email/:email", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}
