package usecase

import (
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"time"
)

type CreateSensorLogUseCase struct {
	SensorRepository entity.SensorRepository
}

type CreateSensorLogInputDTO struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

type CreateSensorLogOutputDTO struct {
	ID        string    `json:"id"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func NewCreateSensorLogUseCase(sensorRepository entity.SensorRepository) *CreateSensorLogUseCase {
	return &CreateSensorLogUseCase{SensorRepository: sensorRepository}
}

func (c *CreateSensorLogUseCase) Execute(input CreateSensorLogInputDTO) (*CreateSensorLogOutputDTO, error) {
	log := entity.NewLog(input.ID, input.Data)
	err := c.SensorRepository.CreateSensorLog(log)
	if err != nil {
		return nil, err
	}
	return &CreateSensorLogOutputDTO{
		ID:        log.ID,
		Data:      log.Data,
		Timestamp: log.Timestamp,
	}, nil
}
