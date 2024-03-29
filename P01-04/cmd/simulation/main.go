package main

import (
	"context"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/repository"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"
	"time"
	"encoding/json"
)

func main() {
	options := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s", os.Getenv("MONGODB_ATLAS_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGODB_ATLAS_CLUSTER_HOSTNAME"), os.Getenv("MONGODB_ATLAS_APP_NAME")))
	client, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewSensorRepositoryMongo(client, "mongodb", "sensors")
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		log.Printf("Starting sensor: %v", sensor)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("ssl://%s:%s", os.Getenv("BROKER_TLS_URL"), os.Getenv("BROKER_PORT"))).SetUsername(os.Getenv("BROKER_USERNAME")).SetPassword(os.Getenv("BROKER_PASSWORD")).SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			for {
				payload, err := entity.NewSensorPayload(
					sensor.ID,
					sensor.Params,
					time.Now(),
				)
				if err != nil {
					log.Fatalf("Failed to create payload: %v", err)
				}

				jsonBytesPayload, err := json.Marshal(payload)
				if err != nil {
					log.Println("Error converting to JSON:", err)
				}
				token := client.Publish("sensors", 1, false, string(jsonBytesPayload))
				log.Printf("Published: %s, on topic: %s", string(jsonBytesPayload), "sensors")
				token.Wait()
				time.Sleep(360 * time.Second)
			}
		}(sensor)
	}
	wg.Wait()
}
