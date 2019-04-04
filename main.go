package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

const topicName = "test"

func main() {
	// The ID of your GCP project
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	t := client.Topic(topicName)
	publishMsg(t)
}

func publishMsg(t *pubsub.Topic) {
	ctx := context.Background()
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello Topic"),
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Unable to publish msg: %s", err.Error())
	}
	log.Printf("Successfully published msg with ID: %s", id)
}
