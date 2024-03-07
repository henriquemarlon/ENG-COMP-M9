// init-mongo.go
package main

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	options := options.Client().ApplyURI("mongodb+srv://development:Lux314159@orbit-city.qa28u1u.mongodb.net/?retryWrites=true&w=majority&appName=orbit-city")
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established successfully")
	db := client.Database("mongodb")
	collection := db.Collection("sensors")
	documents := []interface{}{
		map[string]interface{}{
			"name":      "MICS-6814",
			"latitude":  -23.595354,
			"longitude": -46.664280,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.5618,
			"longitude": -46.6846,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "RXW-LIB-900",
			"latitude":  -23.5894,
			"longitude": -46.6344,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "MICS-6814",
			"latitude":  -23.5788,
			"longitude": -46.6708,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
		map[string]interface{}{
			"name":      "SPS30",
			"latitude":  -23.5718,
			"longitude": -46.708,
			"params": map[string]interface{}{
				"co2":  map[string]interface{}{"min": 0, "max": 1000, "z": 1.96},
				"co":   map[string]interface{}{"min": 0, "max": 15, "z": 1.96},
				"no2":  map[string]interface{}{"min": 0, "max": 1130, "z": 1.96},
				"mp10": map[string]interface{}{"min": 0, "max": 250, "z": 1.96},
				"mp25": map[string]interface{}{"min": 0, "max": 125, "z": 1.96},
				"rad":  map[string]interface{}{"min": 1, "max": 1280, "z": 1.96},
			},
		},
	}
	insertResult, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documents inserted. IDs: %v\n", insertResult.InsertedIDs)
}
