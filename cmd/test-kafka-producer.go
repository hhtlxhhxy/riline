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
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		go func() {
			for {
				producer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder(TestMessage), Partition:int32(i)}
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Hour * 2)
}
