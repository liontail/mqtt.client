package mqttclient

import (
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient *mqtt.Client

func Connect(userName, password, url string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s", url))
	opts.SetUsername(userName)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	token := client.Connect()

	for !token.WaitTimeout(3 * time.Second) {
		return nil, errors.New("Connection MQTT Timeout")
	}
	if err := token.Error(); err != nil {
		return nil, err
	}
	mqttClient = &client
	// fmt.Println("MQTT", url, "Connected")
	return client, nil
}

func GetClient() mqtt.Client {
	return *mqttClient
}

func ListenTo(client mqtt.Client, topic string, f func(mqtt.Message)) {
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		f(msg)
	})
}

func GetMessageFromBeginning(client mqtt.Client, clientName, dbName string) ([]byte, error) {
	serviceName := fmt.Sprintf("%s-%s", clientName, time.Now().Format("150405"))
	data := []byte{}
	forever := make(chan bool)
	var err error
	f := func(m mqtt.Message) {
		go func(ch chan bool) {
			time.Sleep(time.Second * 5)
			err = fmt.Errorf("Cannot Get all data from: %s", dbName)
			<-forever
		}(forever)
		data = m.Payload()
		err = nil
		forever <- true
	}
	go ListenTo(client, serviceName, f)
	token := client.Publish("master", 0, false, fmt.Sprintf(`{ "op": "pull", "clientId": "%s", "dbname": "%s" }`, serviceName, dbName))
	if token.Error() != nil {
		// fmt.Println("Error Publish:", token.Error())
		return nil, token.Error()
	}
	<-forever
	return data, err
}
