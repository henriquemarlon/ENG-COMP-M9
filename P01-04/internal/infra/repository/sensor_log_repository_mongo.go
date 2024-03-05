package repository

import (
	"context"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type SensorLogRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorLogRepositoryMongo(client *mongo.Client, dbName, collection, logsCollection string) *SensorRepositoryMongo {
	sensorsColl := client.Database(dbName).Collection(collection)
	return &SensorRepositoryMongo{
		Collection: sensorsColl,
	}
}

func (s *SensorRepositoryMongo) CreateSensorLog(sensorLog *entity.Log) error {
	_, err := s.Collection.InsertOne(context.TODO(), sensorLog)
	log.Printf("Inserting log into the MongoDB collection")
	return err
}
