package helpers

import (
	"Auth-Golang/database"
	"context"
	"log"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
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

var UserCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastName string, userType string, uid string)(signedtoken string, signedrefreshtoken string){
	claims:= &SingnedDetails{
		Email: email,
		First_name: firstname,
		Last_name: lastName,
		User_Type: userType,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour* time.Duration(24)).Unix(),
		}, 
	}

	refreshclaims:= &SingnedDetails{
      StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour* time.Duration(168)).Unix(),
	  },
	}

	token,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))

	if err!= nil{
		log.Panic(err)
	}

	return token, refreshToken
}

func UpdateAllTokens(signedtoken string , signedrefreshtoken string, userId string){
  var ctx,cancel = context.WithTimeout(context.Background(),100*time.Second)
   var UpdateObj primitive.D

   UpdateObj = append(UpdateObj, bson.E{"Token",signedtoken})
   UpdateObj = append(UpdateObj, bson.E{"RefreshToken",signedrefreshtoken})
   Updated_at ,_ := time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
   UpdateObj = append(UpdateObj, bson.E{"UpdatedAt",Updated_at})

   upsert := true
   filter := bson.M{"userID":userId}

   opt:= options.UpdateOptions{
	Upsert: &upsert,
   }

   _,err:= UserCollection.UpdateOne(
	ctx,
	filter,
	bson.D{
		{"$set",UpdateObj},
	},
	&opt,
   )
   defer cancel()
   if err!= nil{
	log.Panic(err)
	return
   }
   
   return
}

func ValidateToken(signedToken string){


	return

}