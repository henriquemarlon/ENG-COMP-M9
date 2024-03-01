package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/infra/mqtt"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/infra/repository"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/usecase"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/infra/web"
	_ "github.com/lib/pq"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	db, err := sql.Open("postgres", "postgresql://admin:password@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	msgChan := make(chan string)

	sensorsRepository := repository.NewSensorRepositoryPostgres(db)
	createSensorLogUseCase := usecase.NewCreateSensorLogUseCase(sensorsRepository)
	createSensorUseCase := usecase.NewCreateSensorUseCase(sensorsRepository)
	sensorHandlers := web.NewSensorHandlers(createSensorUseCase)

	alertRepository := repository.NewAlertRepositoryPostgres(db)
	createAlertUseCase := usecase.NewCreateAlertUseCase(alertRepository)
	findAllAlertsUseCase := usecase.NewFindAllAlertsUseCase(alertRepository)
	alertHandlers := web.NewAlertHandlers(createAlertUseCase, findAllAlertsUseCase)

	//TODO: this is the best way to do this? need to refactor or find another way to start the server
	r := chi.NewRouter()
	r.Get("/alerts", alertHandlers.FindAllAlertsHandler)
	r.Post("/alerts", alertHandlers.CreateAlertHandler)
	r.Post("/sensors", sensorHandlers.CreateSensorHandler)
	go http.ListenAndServe(":8080", r)

	for msg := range msgChan {
		dto := usecase.CreateSensorLogInputDTO{}
		err := json.Unmarshal(msg, &dto)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		_, err = createSensorLogUseCase.Execute(dto)
		if err != nil {
			log.Fatalf("Failed to create sensor log: %v", err)
		}
	}
}