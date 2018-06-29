package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

	topic := "test"
	topicInfo := "{\"timestamp\":1529482188179,\"action\":\"device.info\",\"topic\":\"device_info\",\"net\":\"MOZI_DEV\",\"cardAvailable\":0,\"cardTotal\":6,\"electricity\":101,\"volume\":25,\"mac\":\"10a4bea634a6\",\"ip\":\"113.89.99.205\",\"powerState\":1,\"earLightStatus\":0,\"childLockStatus\":0,\"firmwareVersion\":\"1.1.10.180525\",\"int_key\":0,\"string_key\":\"test\"}"

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(topicInfo)}
		select {
		case producer.Input() <- message:
			fmt.Printf("%v\r\n", message)
			time.Sleep(2000 * time.Millisecond)
			enqueued++
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			return
			break ProducerLoop
		}
	}
	wg.Wait()
	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}
