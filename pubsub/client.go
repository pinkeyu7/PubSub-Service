package pubsub

import (
	"context"
	"fmt"
	"log"
	"pubsub-event-bus/protobuf/pubsub"

	"google.golang.org/grpc"
)

// PubSubClient encapsulates the gRPC client and connection
type PubSubClient struct {
	conn   *grpc.ClientConn
	client pubsub.PubSubClient
}

// NewPubSubClient initializes a new PubSubClient
func NewPubSubClient(serverAddr string) (*PubSubClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}
	client := pubsub.NewPubSubClient(conn)
	return &PubSubClient{conn: conn, client: client}, nil
}

// Close closes the gRPC connection
func (p *PubSubClient) Close() {
	if p.conn != nil {
		p.conn.Close()
	}
}

// Subscribe listens to a topic and processes messages using a callback function
func (p *PubSubClient) Subscribe(ctx context.Context, topic, subscription string, callback func(message string)) error {
	req := &pubsub.SubscribeRequest{
		Topic:        topic,
		Subscription: subscription,
	}

	stream, err := p.client.Subscribe(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	log.Printf("Subscribed to topic '%s'. Waiting for messages...", topic)

	for {
		msg, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("error receiving message: %w", err)
		}
		callback(msg.Message)
	}
}

// Publish sends a message to the specified topic
func (p *PubSubClient) Publish(ctx context.Context, topic, message string) error {
	req := &pubsub.PublishRequest{
		Topic:   topic,
		Message: message,
	}
	_, err := p.client.Publish(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	return nil
}
