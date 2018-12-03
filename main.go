package mqttclient

import (
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient *mqtt.Client

func Connect(userName, password, url string) error {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", url))
	opts.SetUsername(userName)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {
		return errors.New("Connection MQTT Timeout")
	}
	if err := token.Error(); err != nil {
		return err
	}
	mqttClient = &client
	fmt.Printf("MQTT %s Connected", url)
	return nil
}

func GetClient() *mqtt.Client {
	return mqttClient
}

func ListenTo(client *mqtt.Client, topic string, f func(mqtt.Message)) {
	(*client).Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		// fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		f(msg)
	})
}
