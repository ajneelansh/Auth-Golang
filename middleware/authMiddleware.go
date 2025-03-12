package middleware

import (
	"net/http"
    "Auth-Golang/helpers"
	"gitub.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
     return func(c *gin.context){
		clientToken ,err:= c.Request.Header.Get("token")
		
        if clientToken == ""{ 
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			c.Abort()
			return
		}

		claims,err:= helpers.ValidateToken(clientToken)

		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}

			c.JSON(http.StatusOK)
 }
}
