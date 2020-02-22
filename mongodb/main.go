package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"time"
)

type Person struct {
	ID primitive.ObjectID `json: "_id, omitempty" bson:"_id, omitempty"`
	Firstname string `json:"firstname, omitempty" bson: "_id, omitempty"`
	Lastname string `json:"lastname, omitempty" bson: "lastname, omitempty"`
}

var client *mongo.Client

func main()  {
	fmt.Println("Starting the aplicacion")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, "mongodb://localhost:27017")
	router := mux.NewRouter()
	http.ListenAndServe(":12345", router)
}
