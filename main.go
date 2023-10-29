package main

import (
	"REST-API/controllers"
	"REST-API/services"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	userservice services.UserService
	// controller     controllers.UserController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoConnect   *mongo.Client
	err            error
	usercontoller  controllers.UserController
)

func init() {
		err := godotenv.Load()

	ctx = context.TODO()

	MONGO_URL:= os.Getenv("MONGO_URL")

	mongoOptions := options.Client().ApplyURI(MONGO_URL)
	mongoConnect, err = mongo.Connect(ctx, mongoOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoConnect.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MongoDB connection Established")

	userCollection = mongoConnect.Database("go-rest-api").Collection("users")
	userservice = services.NewUserService(userCollection, ctx)
	usercontoller = controllers.New(userservice)

	server = gin.Default()

}
func main() {
	defer mongoConnect.Disconnect(ctx)

	basePath := server.Group("/api/v1")
	usercontoller.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":3000"))
}
