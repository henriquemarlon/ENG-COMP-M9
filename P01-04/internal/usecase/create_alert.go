package usecase

import (
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CreateAlertUseCase struct {
	AlertRepository entity.AlertRepository
}

type CreateAlertInputDTO struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Option    string  `json:"option"`
}

type CreateAlertOutputDTO struct {
	ID        primitive.ObjectID `json:"_id"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Option    string             `json:"option"`
	Timestamp time.Time          `json:"timestamp"`
}

func NewCreateAlertUseCase(alertRepository entity.AlertRepository) *CreateAlertUseCase {
	return &CreateAlertUseCase{AlertRepository: alertRepository}
}

func (c *CreateAlertUseCase) Execute(input CreateAlertInputDTO) (*CreateAlertOutputDTO, error) {
	alert := entity.NewAlert(input.Latitude, input.Longitude, input.Option)
	id, err := c.AlertRepository.CreateAlert(alert)
	if err != nil {
		return nil, err
	}
	return &CreateAlertOutputDTO{
		ID:        id.InsertedID.(primitive.ObjectID),
		Latitude:  alert.Latitude,
		Longitude: alert.Longitude,
		Option:    alert.Option,
		Timestamp: alert.Timestamp,
	}, nil
}
