package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jvpr123/golang-crud/src/controller"
)

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/getUserById/:userId", controller.FindUserById)
	rg.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	rg.POST("/createUser", controller.CreateUser)
	rg.PUT("/updateUser/:userId", controller.UpdateUser)
	rg.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
