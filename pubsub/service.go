package pubsub

import (
	"context"
	"log"
	"pubsub-event-bus/protobuf/pubsub"
)

type PubSubService struct {
	pubsub.UnimplementedPubSubServer

	context context.Context
	queue   PubSub
}

func NewPubSubService() (*PubSubService, func()) {
	ctx, cancel := context.WithCancel(context.Background())

	queue := New()

	// Initialize the queue
	if err := queue.Init(ctx); err != nil {
		log.Fatalf("Failed to initialize queue: %v", err)
	}

	// defer
	go func() {
		<-ctx.Done()
		if err := queue.Close(); err != nil {
			log.Printf("Failed to close queue: %v", err)
		}
		log.Println("Queue closed")
	}()

	return &PubSubService{
		context: ctx,
		queue:   queue,
	}, cancel
}

func (s *PubSubService) Subscribe(req *pubsub.SubscribeRequest, stream pubsub.PubSub_SubscribeServer) error {
	ctx := stream.Context()
	err := s.queue.Subscribe(ctx, &SubscribeRequest{Topic: req.Topic}, func(ctx context.Context, msg *NewMessage) error {
		response := &pubsub.SubscribeResponse{
			Topic:        msg.Topic,
			Subscription: req.Subscription,
			Message:      string(msg.Data),
		}
		if err := stream.Send(response); err != nil {
			log.Printf("Failed to send message to stream: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("Failed to subscribe to topic: %v", err)
		return err
	}

	select {
	case <-s.context.Done():
		log.Println("Service context done, unsubscribing...")
	case <-ctx.Done():
		log.Println("Context done, unsubscribing...")
	}

	return nil
}

func (s *PubSubService) Publish(ctx context.Context, req *pubsub.PublishRequest) (*pubsub.PublishResponse, error) {
	err := s.queue.Publish(ctx, &PublishRequest{
		Topic: req.Topic,
		Data:  []byte(req.Message),
	})
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return nil, err
	}
	return &pubsub.PublishResponse{
		Topic:   req.Topic,
		Message: req.Message,
	}, nil
}
