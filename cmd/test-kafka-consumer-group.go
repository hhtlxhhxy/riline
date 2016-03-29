package main

import (
	"github.com/Shopify/sarama"
	"os"
	"fmt"
	"time"
)

const TestMessage = "ABC THE MESSAGE"

func main() {
	config := sarama.NewConfig()
	config.Producer.Flush.Messages = 100
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go func() {
		for {
			producer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder(TestMessage)}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	master, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// When
	consumer, err := master.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Then: messages starting from offset 1234 are consumed.
	for {
		select {
		case message := <-consumer.Messages():
			fmt.Println(message)
		case err := <-consumer.Errors():
			fmt.Println(err)
			os.Exit(1)
		}
	}

}