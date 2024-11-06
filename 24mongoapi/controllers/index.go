package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/enigma0844/mongoapi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const colName = "watchlist"
const dbName = "netflix"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection referenced")
}

// MongoDB helpers

func insertOneMovie(movie models.Netflix) {
	doc, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie with ID:", doc.InsertedID)
}

func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"watched": true}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modifed count:", doc.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", doc.DeletedCount)
}

func deleteAllMovies() {
	doc, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bulk deleted count:", doc.DeletedCount)
}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

// actual controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"] + " Successfully updated to watched")
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"] + " successfully deleted")
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies deleted")
}
