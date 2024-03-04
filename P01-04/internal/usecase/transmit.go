package usecase

import (
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
)

type TransmitUseCase struct {
	LogRepository entity.LogRepository
}

func NewTransmitUseCase(logRepository entity.LogRepository) *TransmitUseCase {
	return &TransmitUseCase{LogRepository: logRepository}
}

func (t *TransmitUseCase) Execute(log *entity.Log) error {
	err := t.LogRepository.Transmit(log)
	if err != nil {
		return err
	}
	return nil
}