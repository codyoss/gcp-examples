package main

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
)

const subscriberName = "test-sub"

func main() {
	// The ID of your GCP project
	projectID := os.Getenv("GCP_PROJECT_ID")

	// Creates a client to GCP pubsub
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Creates the subscriber(aka consumer)
	sub := client.Subscription(subscriberName)

	// Receive a message. Passing in a cancel context is done to controll all the goroutines this method will spawn
	// under the hood. Receive blocks the current thread, so launched it out to its own goroutine.
	cctx, cancel := context.WithCancel(ctx)
	go sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("ID: %s\tData: %s", msg.ID, string(msg.Data))
		// Tells pubsub you successfully consumed this message so it will not be sent again.
		msg.Ack()
	})

	// Give the app a second to consume a message
	time.Sleep(3 * time.Second)

	// This will stop all the goroutines that `sub.Receive` spawns under the hood
	cancel()

	client.Close()
}
