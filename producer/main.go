package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {

	w := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "example-123",
		RequiredAcks:           kafka.RequireAll,
		AllowAutoTopicCreation: true,
		Async:                  true,
		Completion: func(messages []kafka.Message, err error) {

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, val := range messages {
				fmt.Printf("messages sent, offset %d, key %s, val %s \n", val.Offset, val.Key, val.Value)
			}
		},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("yash"),
			Value: []byte("gautam"),
		},
		kafka.Message{
			Key:   []byte("anubhav"),
			Value: []byte("prasad"),
		},
		kafka.Message{
			Key:   []byte("arya"),
			Value: []byte("vashishth"),
		},
		kafka.Message{
			Key:   []byte("himanshu"),
			Value: []byte("atri"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
