package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/rabbitmq"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/repository"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/infra/web"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/usecase"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USERNAME"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	msgChan := make(chan *amqp.Delivery)
	rabbitmqRepository := rabbitmq.NewRabbitMQConsumer("sensors-log-queue", "amqp://app:admin1234@rabbitmq:5672/")
	go rabbitmqRepository.Consume(msgChan)

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
		err := json.Unmarshal(msg.Body, &dto)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}
		_, err = createSensorLogUseCase.Execute(dto)
		if err != nil {
			log.Fatalf("Failed to create sensor log: %v", err)
		}
	}
}
