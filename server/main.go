package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	pubsubProto "pubsub-event-bus/protobuf/pubsub"
	"pubsub-event-bus/pubsub"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	pubSubService, cancelFunc := pubsub.NewPubSubService()

	// Create the gRPC server
	grpcServer := grpc.NewServer()
	pubsubProto.RegisterPubSubServer(grpcServer, pubSubService)

	// Start listening on a port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down gRPC server...")
		cancelFunc()
		grpcServer.GracefulStop()
	}()

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
