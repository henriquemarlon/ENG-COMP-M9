package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/henriquemarlon/ENG-COMP-M9/P01-04/internal/domain/entity"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type SubscriberMQTTRepository struct {
	Client MQTT.Client
}

func NewSubscriberMQTTRepository(client MQTT.Client) *SubscriberMQTTRepository {
	return &SubscriberMQTTRepository{
		Client: client,
	}
}

func (s *SubscriberMQTTRepository) Subscribe(data *entity.Log) error {
	for {
		token := s.Client.Subscribe("/stations", 1, func(client MQTT.Client, message MQTT.Message) {
			var payload entity.Log
			err := json.Unmarshal(message.Payload(), &payload)
			if err != nil {
				fmt.Println("Error converting to JSON:", err)
				return
			}
			*data = payload
		})
		token.Wait()
	}
}
