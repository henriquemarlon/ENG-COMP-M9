package station

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math"
	"math/rand"
	"time"
	"github.com/henriquemarlon/ENG-COMP-M9/ativ01"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Station struct {
	Location string `json:"sensor"`
	Gas      string `json:"gas"`
	Rad_Lum  string `json:"rad_lum"`
}

type Interval struct {
	Minimum float64
	Maximum float64
}

var Area = map[string]Interval{
	"latitude":  {0, 39},
	"longitude": {0, 39},
}

func LocationEntropy(key string) float64 {
	rand.NewSource(time.Now().UnixNano())
	max := Area[key].Maximum
	min := Area[key].Minimum
	value := rand.Float64()*(max-min) + min
	return math.Round(value)
}

func GenerateLocation() Location {
	data := Location{
		Latitude:  LocationEntropy("latitude"),
		Longitude: LocationEntropy("longitude"),
	}
	return data
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID(fmt.Sprintf("station-%d", rand.Intn(1000)))
	client := MQTT.NewClient(opts)
	if session := client.Connect(); session.Wait() && session.Error() != nil {
		panic(session.Error())
	}


	for {
		// Cria a estrutura de dados para enviar ao broker MQTT
		senddata := SendData{
			Sensor:           config.Sensor,
			Latitude:         config.Latitude,
			Longitude:        config.Longitude,
			QoS:              config.QoS,
			Unit:             config.Unit,
			TransmissionRate: config.TransmissionRate,
			CurrentTime:      time.Now(),
			Values:           createGasesValues(),
		}

		jsonData, err := json.MarshalIndent(senddata, "", "    ")
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
			return
		}

		token := client.Publish(config.Sensor, config.QoS, false, string(jsonData))
		token.Wait()
		fmt.Println("Publicado:", string(jsonData))
		time.Sleep(time.Duration(config.TransmissionRate) * time.Second)
	}
}
