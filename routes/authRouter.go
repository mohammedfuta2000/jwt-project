package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedfuta2000/jwt-project/controllers"
)

func AuthRoutes(incommingRoutes *gin.Engine)  {
	incommingRoutes.POST("/users/signup", controllers.Signup())
	incommingRoutes.POST("/users/login", controllers.Login())
}