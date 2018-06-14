package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// callback function
var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	emqtt()
}

func emqtt() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	//connect mqtt-server and set clientID
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("mqtt_server")

	//set userName
	opts.SetUsername("mqtt_server")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	// create object
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// publish topic
	for i := 0; i < 180; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/sample", 0, false, text)
		time.Sleep(1 * time.Second)
		token.Wait()
	}

	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}
