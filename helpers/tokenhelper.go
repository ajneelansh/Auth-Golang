package helpers

import(
	"fmt"
	"os"
	"time"
	"context"
	"log"
	"Auth-Golang/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SingnedDetails struct{
  Email string
  First_name string
  Last_name string 
  Uid string
  User_Type string
  jwt.StandardClaims
}

var UserCollection * mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastName string, userType string, uid string)(signedtoken string, signedrefreshtoken string){
	calims:= &SingnedDetails{
		Email: email,
		First_name: firstname,
		Last_Name: lastName,
		User_Type: userType,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour* time.Duration(24)).Unix(),
		}, 
	}

	refreshclaims:= &SingnedDetails{
      StandardClaims: jwt.StandardClaims{
		ExpiresAt: time,Now().Local().Add(time.Hour* time.Duration(168)).Unix()
	  }
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))

	if err!=nil{
		log.Panic(err)  
		return  
	}

	return token, refreshToken ,err
}