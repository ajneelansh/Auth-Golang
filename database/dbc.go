package database

import(
	"fmt"
	"os"
	"log"
	"time"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
  err := godotenv.Load(".env")
  if err != nil{
	log.Fatal("Error Loading .env File")
  }

  MongoDb := os.Getenv("MONGODB_URL")
   opt := options.Client().ApplyURI(MongoDb)

  ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  client,err := mongo.Connect(ctx,opt)
  if err!=nil{
	log.Fatal(err)
  }

  fmt.Println("Connected to MongoDb")

  return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string ) * mongo.Collection{
	var collection  *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}