package databaseservice

import (
	"context"
	"log"
	"time"
	//"net/http"
	//"go/types"
	//"encoding/json"
	//Structs
	databasestructs "malikiah.io/structs/databaseStructs"
	
	//Services
	//passwordservice "malikiah.io/services/passwordService"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"

)

func Find(databaseQuery databasestructs.Find) (user databasestructs.User, page string) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Println(err)
		}
	// Database name
	db := client.Database("malikiah")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	
	//Mongo Collection to search
	collection := db.Collection(databaseQuery.MongoCollection)

	// Search Filter
	filter := bson.M{databaseQuery.Criteria: databaseQuery.CriteriaValue}

	if databaseQuery.Criteria == "_id"{
		filter = bson.M{databaseQuery.Criteria: databaseQuery.ID}
	} else if databaseQuery.Criteria == "" { 
		// Finds entire collection

	} else if databaseQuery.FindAll == true {
		// Finds all that contain a specific Criteria and Criteria Value pair
		filter = bson.M{}
	} else {
		filter = bson.M{databaseQuery.Criteria: databaseQuery.CriteriaValue}
	}

	switch databaseQuery.MongoCollection {

	case "user":
		if databaseQuery.Criteria == "" {
				//collection.Find(ctx, filter).Decode(&page)
			} else { 
				collection.FindOne(ctx, filter).Decode(&user)
		 	}
	case "page":
		if databaseQuery.Criteria == "" {
				//collection.Find(ctx, filter).Decode(&page)
			} else { 
				collection.FindOne(ctx, filter).Decode(&page)
			 }
	}

		
		
	return
}

func Insert(mongoCollection string, data databasestructs.User) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	
	// Database name
	db := client.Database("malikiah")

	log.Println(data)
	collection := db.Collection(mongoCollection)

	insertResult, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(insertResult.InsertedID)
}