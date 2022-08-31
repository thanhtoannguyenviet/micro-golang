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
	v1 := r.Group("/v1/auth")
	v1.POST("regist", user.InsertUser(client))
	r.Run()
	//database, err := client.ListDatabases(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(database)
	//collection := client.Database("micro-go").Collection("auth-user")
	//cur, err := collection.Find(ctx, bson.D{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cur.Close(ctx)
	//var elems []user.UserModel
	//if err = cur.All(ctx, &elems); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(elems)
}
