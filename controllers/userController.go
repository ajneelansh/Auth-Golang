package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/ajneelansh/Auth-Golang/database"
	"github.com/ajneelansh/Auth-Golang/helpers"
	"github.com/ajneelansh/Auth-Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.new()

func HashPasswaord()

func VerifyPassword()

func Signup()

func Login()

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

