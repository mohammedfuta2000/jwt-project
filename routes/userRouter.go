package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohammedfuta2000/jwt-project/controllers"
	"github.com/mohammedfuta2000/jwt-project/middleware"
)

func UserRoutes(incommingRoutes *gin.Engine)  {
	// to get users, just like u get products
	incommingRoutes.Use(middleware.Authenticate())
	// all routes after this line will require authentications henceforth
	incommingRoutes.GET("/users",controllers.GetUsers())
	incommingRoutes.GET("/users/:user_id",controllers.GetUser())
}