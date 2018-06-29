package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	cluster "github.com/bsm/sarama-cluster"
)

func main() {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"kafka:9092"}
	topics := []string{"test"}
	consumer, err := cluster.NewConsumer(brokers, "test-group", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				go consumePoint(msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}

type filter struct {
	Url    string
	Topic  string
	Action string
	Point  string
	Filter int
	Triger string
}

type msgInfo struct {
	Action    string
	Topic     string
	Timestamp int64
}

func consumePoint(msg []byte) {
	message := msgInfo{}
	err := json.Unmarshal(msg, &message)
	if err != nil {
		fmt.Printf("%v\r\n", err)
	}
	fmt.Printf("%v\r\n", message)
}
