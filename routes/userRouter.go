package routes

import(
  "Auth-Golang/controllers"
  "Auth-Golang/middleware"
  "github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
  incomingRoutes.Use(middleware.Authenticate())
  incomingRoutes.POST("users", controllers.GetUsers())
	incomingRoutes.POST("users/:user_id", controllers.GetUser())  
}