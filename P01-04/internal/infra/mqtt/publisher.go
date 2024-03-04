package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type PublisherMQTTRepository struct {
	Client MQTT.Client
}

func NewPublisherMQTTRepository(client MQTT.Client) *PublisherMQTTRepository {
	return &PublisherMQTTRepository{
		Client: client,
	}
}

func (s *PublisherMQTTRepository) Publish(data *entity.Log) error {
	for {
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return err
		}
		token := s.Client.Publish("sensors", 1, false, string(payload))
		log.Println("Published: ", string(payload))
		token.Wait()
		time.Sleep(3 * time.Second)
	}
}
