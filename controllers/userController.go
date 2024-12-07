package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"Auth-Golang/database"
	"Auth-Golang/helpers"
	"Auth-Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()

func HashPasswaord()

func VerifyPassword(userPassword string, ProvidedPassword string){
  bcrypt.CompareHashAndPassword()
}

func Signup() gin.HandlerFunc{
	return func(c *gin.Context){

	ctx,cancel := context.WithTimeout(context.Background(),100*time.Second) 
    var user models.User

	if err := c.BindJSON(&user); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	count, err := UserCollection.CountDocuments(ctx, bson.M{"email":user.Email})
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		log.Panic(err)
	}

	if count>0{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"This email already exists"})
	}

	user.CreatedAt,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID= primitive.NewObjectID()
	user.UserID = user.ID.Hex()

	token,refreshToken:= helpers.GenerateAllTokens(user.Email,user.FirstName,user.LastName,user.Role,user.UserID)

	user.Token = token
	user.RefreshToken = refreshToken

	resultInsertionNo , inserterr := UserCollection.InsertOne(ctx,user)
	if inserterr!= nil{
		msg:= fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
		return
	}

	defer cancel()
	c.JSON(http.StatusOK , resultInsertionNo)


 }
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx,cancel:= context.WithTimeout(context.Background(), 100*time.Second)
        var user models.User
		var founduser models.User

		if err:= c.BindJSON(&user); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		err:= UserCollection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"email or password is incorrect"})
			return
		}

		VerifyPassword(user.Password,founduser.Password)

	}
}

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		userId:= c.Param("User_id")

		if err := helpers.MatchUserTypeToId(c,userId); err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"errror":err.Error()})
			return
		}

		ctx,cancel := context.WithTimeout(context.Background(), 100*time.Second)

	    var user models.User
		err :=UserCollection.FindOne(ctx, bson.M{"id":userId}).Decode(&user)

		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, user)
	}
}

