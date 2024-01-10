package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/kiyoshi-87/RetireMate-GO-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const dbName string = "RetireMate-Users"

const collectionName string = "user-info"

func init() {

	// Construct the full path to the .env file

	if err := godotenv.Load("go.env"); err != nil {
		log.Println("No .env file found!")
	}

	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB is succesfully connected!")
	collection = client.Database(dbName).Collection(collectionName)
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API RUNNING SUCCESFULLY!"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Userdata
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	result, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	email := params["email"]

	filter := bson.M{"email": email}
	var data bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&data) //All the data is embeded in the form of bson inside the variable data

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(data)
}

func UpdateUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	email := params["email"]

	var data models.Userdata
	json.NewDecoder(r.Body).Decode(&data)
	filter := bson.M{"email": email}
	update := bson.M{"$set": data}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
