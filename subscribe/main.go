package main

import (
	"context"
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

	// Subscribe to the topic "main" and process messages
	err = pubSubClient.Subscribe(context.Background(), "main", "subscriber-1", handler)
	if err != nil {
		log.Fatalf("Subscription error: %v", err)
	}
}

func handler(message string) {
	log.Printf("Received message: %s", message)
	time.Sleep(1 * time.Second) // Simulate processing delay
}
