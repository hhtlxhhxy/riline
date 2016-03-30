package main

import (
	"github.com/Shopify/sarama"
	"os"
	"fmt"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Flush.Messages = 100

	for i := 0; i < 3; i++ {
		master, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		consumer, err := master.ConsumePartition("my_topic", int32(i), sarama.OffsetNewest)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go func() {
			for {
				select {
				case message := <-consumer.Messages():
					fmt.Println("Key:", message.Key,
						"Value:", string(message.Value),
						"Partition:", message.Partition,
						"Offset:", message.Offset)
				case err := <-consumer.Errors():
					fmt.Println(err)
					os.Exit(1)
				}
			}
		}()
	}
	time.Sleep(1 * time.Hour)
}