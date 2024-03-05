package repository

import (
	"context"
	"log"

	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SensorRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewSensorRepositoryMongo(client *mongo.Client, dbName, sensorsCollection, logsCollection string) *SensorRepositoryMongo {
	sensorsColl := client.Database(dbName).Collection(sensorsCollection)
	return &SensorRepositoryMongo{
		Collection: sensorsColl,
	}
}

func (s *SensorRepositoryMongo) CreateSensor(sensor *entity.Sensor) error {
	_, err := s.Collection.InsertOne(context.TODO(), sensor)
	log.Printf("Inserting sensor into the MongoDB collection")
	return err
}

func (s *SensorRepositoryMongo) FindAllSensors() ([]*entity.Sensor, error) {
	cur, err := s.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var sensors []*entity.Sensor
	for cur.Next(context.TODO()) {
		var sensor entity.Sensor
		err := cur.Decode(&sensor)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return sensors, nil
}
