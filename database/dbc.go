package database

import(
	"fmt"
	"os"
	"log"
	"time"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go/mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client{
  err:= godotenv.Load(".env")
  if err != nil{
	log.Fatal("Error Loading .env File")
  }

  mongoDb := os.Getenv("MONGODB_URL")
}