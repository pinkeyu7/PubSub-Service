package main

import (
	"context"
	"fmt"
	"log"
	"pubsub-event-bus/pubsub"
	"time"
)

func main() {
	// Initialize the PubSubClient
	pubSubClient, err := pubsub.NewPubSubClient("localhost:50051")
	if err != nil {
		log.Fatalf("Error initializing PubSubClient: %v", err)
	}
	defer pubSubClient.Close()

	// Publish 10 messages to the topic "main"
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		err := pubSubClient.Publish(context.Background(), "main", message)
		if err != nil {
			log.Printf("Failed to publish message %d: %v", i, err)
			continue
		}
		log.Printf("Published message %d: %s", i, message)
		time.Sleep(1 * time.Second) // Simulate delay between messages
	}
}
