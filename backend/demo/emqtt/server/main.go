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
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("4000100020003000")

	//set userName
	opts.SetUsername("OGY0ZTRiMzVjMGNi")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	// create object
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	msg := "{\"timestamp\":1529482188179,\"action\":\"device.info\",\"topic\":\"device_info\",\"net\":\"MOZI_DEV\",\"cardAvailable\":0,\"cardTotal\":6,\"electricity\":101,\"volume\":25,\"mac\":\"10a4bea634a6\",\"ip\":\"113.89.99.205\",\"powerState\":1,\"earLightStatus\":0,\"childLockStatus\":0,\"firmwareVersion\":\"1.1.10.180525\",\"int_key\":0,\"string_key\":\"test\"}"
	topic := "/OGY0ZTRiMzVjMGNi/clients/4000106300000017/event/device_info"

	// publish topic
	for i := 0; i < 100; i++ {
		token := c.Publish(topic, 0, false, msg)
		time.Sleep(100 * time.Millisecond)
		token.Wait()
	}

	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}
