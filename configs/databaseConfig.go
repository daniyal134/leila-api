package configs

import (
	"context"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"time"
)

	type DBCredentials struct {
		DBHost		string
		DBUsername	string
		Timeout  	time.Duration
	}

	func StartDB() {
		dbCredentials := getDBCredentials()
		client, err := mongo.NewClient(dbCredentials.DBHost)
		if err != nil {
			log.Fatalf("todo: database configuration failed: %v", err)
		} else {
			log.Println("succesfuly connected to mongo")

		}
		ctx, _ := context.WithTimeout(context.Background(), dbCredentials.Timeout)
		err = client.Connect(ctx)
		collection := client.Database("testing").Collection("numbers")
		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		if err != nil {
			log.Fatalf("insertion failed", err)
		} else {
			log.Println("Successfully inserted")
		}
		id := res.InsertedID
		println(id)

	}

	func getDBCredentials() DBCredentials {

		var dbCredentials DBCredentials
		err := envconfig.Process("LEILAAPI", &dbCredentials)
		if err != nil {
		log.Fatal(err)
		}
		fmt.Printf("%+v\n", dbCredentials)
		return dbCredentials
	}

