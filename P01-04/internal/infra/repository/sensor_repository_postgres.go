package repository

import (
	"database/sql"
	"github.com/Inteli-College/2024-T0002-EC09-G04/internal/domain/entity"
	_ "github.com/lib/pq"
)

type SensorRepositoryPostgres struct {
	DB *sql.DB
}

func NewSensorRepositoryPostgres(db *sql.DB) *SensorRepositoryPostgres {
	return &SensorRepositoryPostgres{
		DB: db,
	}
}

func (s *SensorRepositoryPostgres) CreateSensor(sensor *entity.Sensor) error {
	_, err := s.DB.Exec("INSERT INTO sensors (id, name, latitude, longitude) VALUES ($1, $2, $3, $4)", sensor.ID, sensor.Name, sensor.Latitude, sensor.Longitude)
	if err != nil {
		return err
	}
	return nil
}

func (s *SensorRepositoryPostgres) CreateSensorLog(log *entity.Log) error {
	_, err := s.DB.Exec("INSERT INTO sensors_log (sensor_id, data, timestamp) VALUES ($1, $2, $3)", log.ID, log.Data, log.Timestamp)
	if err != nil {
		return err
	}
	return nil
}

func (s *SensorRepositoryPostgres) FindAllSensors() ([]*entity.Sensor, error) {
	rows, err := s.DB.Query("SELECT id, name, latitude, longitude FROM sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sensors []*entity.Sensor
	for rows.Next() {
		var sensor entity.Sensor
		if err := rows.Scan(&sensor.ID, &sensor.Name, &sensor.Latitude, &sensor.Longitude); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}
