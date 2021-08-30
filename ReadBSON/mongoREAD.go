package main

import (
	"context"
	//"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	// Set client options
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Luckyshukla:Luckyshukla@cluster0.mlkjv.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	quickstartDatabase := client.Database("quickstart")
	//podcastsCollection := quickstartDatabase.Collection("podcasts")
	episodesCollection := quickstartDatabase.Collection("episodes")
	// Retrive all  data from databases

	cursor, err := episodesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	// find All data it could be in unorderd

	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}

	// Retrive the data in batches

	/*
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var episode bson.M
			if err = cursor.Decode(&episode); err != nil {
				log.Fatal(err)
			}
			fmt.Println(episode)

		}
	*/

	/*
		var podcast bson.M
		if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
			log.Fatal(err)
		}
		fmt.Println(podcast)
	*/

	//Querying Documents from a Collection with a Filter

	/*
		filterCursor, err := episodesCollection.Find(ctx, bson.M{"duration": 25})
		if err != nil {
			log.Fatal(err)
		}
		var episodefiltered []bson.M
		if err = filterCursor.All(ctx, &episodefiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episodefiltered)
	*/
}
