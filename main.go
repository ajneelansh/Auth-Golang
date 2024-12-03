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

  router.GET("/api-1", func(c *gin.Context){
      c.JSON(200 ,gin.H{"success":"Access to API-1"})
  })
  router.GET("/api-2,",func(c *gin.Context)){
    c.JSON(200,fin.H{"success":"Access to API-2"})
  }

}

