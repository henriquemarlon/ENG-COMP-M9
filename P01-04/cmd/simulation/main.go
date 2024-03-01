package main

import (
	"database/sql"
	"log"
	"sync"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/domain/entity"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/infra/mqtt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/usecase"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://admin:password@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to Database: %v", err)
	}
	defer db.Close()

	repository := repository.NewSensorRepositoryPostgres(db)
	findAllSensorsUseCase := usecase.NewFindAllSensorsUseCase(repository)

	sensors, err := findAllSensorsUseCase.Execute()
	if err != nil {
		log.Fatalf("Failed to find all sensors: %v", err)
	}

	var wg sync.WaitGroup
	for _, sensor := range sensors {
		wg.Add(1)
		go func(sensor usecase.FindAllSensorsOutputDTO) {
			defer wg.Done()
			// opts := MQTT.NewClientOptions().AddBroker("rabbitmq:1883").SetUsername("simulation").SetPassword("admin123456").SetClientID(sensor.ID)
			opts := MQTT.NewClientOptions().AddBroker("rabbitmq:1883").SetUsername("simulation").SetPassword("admin123456").SetClientID(sensor.ID)
			client := MQTT.NewClient(opts)
			if session := client.Connect(); session.Wait() && session.Error() != nil {
				log.Fatalf("Failed to connect to MQTT broker: %v", session.Error())
			}
			stationRepository := mqtt.NewPublisherMQTTRepository(client)
			id, value := entity.NewSensorPayload(
				sensor.ID,
				map[string][]float64{"co2": {0, 100, 3}, "co": {0, 100, 3}, "no2": {0, 100, 3}, "mp10": {0, 100, 3}, "mp25": {0, 100, 3}, "rad": {0, 100, 3}},
			)
			log := entity.NewLog(id, value)
			stationRepository.Publish(log)
		}(sensor)
	}
	wg.Wait()
}
