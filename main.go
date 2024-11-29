package main

import(
 routes "github.com/ajneelansh/Auth-Golang/routes"
 "os"
 "github.com/gin-gonic/gin"

)

func main(){
 port := os.Getenv("PORT")
  if port==""{
	port = "8000"
  }
  router:=gin.New()
  router.Use(gin.Logger())

  routes.AuthRoutes(router)
  routes.UserRoutes(router)

}

