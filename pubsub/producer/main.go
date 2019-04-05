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
	projectID := os.Getenv("GCP_PROJECT_ID")

	// Creates a client to GCP pubsub
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Creates a reference to the topic we want to interact with
	t := client.Topic(topicName)

	// Publishes a message on a topic
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Hello Topic"),
	})

	// blocking call that waits for the messaged to be published or returns an error
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Unable to publish msg: %s", err.Error())
	}
	log.Printf("Successfully published msg with ID: %s", id)

	client.Close()
}
