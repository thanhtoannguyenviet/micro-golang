package main

import (
	"auth/modules/user"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	runServer()
}
func runServer() {
	domain := os.Getenv("domain_mongo")
	if domain == "" {
		domain = "cluster0.sfd8i.mongodb.net"
	}
	username := os.Getenv("username_mongo")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("password_mongo")
	if password == "" {
		password = "root"
	}
	mongo_uri := "mongodb+srv://" + username + ":" + password + "@" + domain + "/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	r := gin.Default()
	auth := r.Group("v1/auth")
	{
		auth.POST("register", user.InsertUser(client))
		auth.POST("login", user.Login(client))
		auth.GET(":id", user.GetById(client))
		auth.GET("", user.GetAll(client))
		auth.PATCH("", user.UpdateUser(client))
	}
	r.Run()
}
