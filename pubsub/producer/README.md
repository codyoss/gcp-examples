# gcloud pubsub producer

Produce messages `pubsub` in 5 easy to follow steps!

The example shows how to produce messages to a `pubsub` topic from Go and consume them from your console with `gcloud`.

## Directions

1. Create a pubsub topic. This is what you will send and recieve data from: `gcloud pubsub topics create test`

2. Create a pubsub subscriber. This is what you will use to pull message from the command line:
`gcloud pubsub subscriptions create test-sub --topic test`

3. Run the provided code: `go run main.go`

4. Pull the data your app just produced! : `gcloud pubsub subscriptions pull test-sub`

5. Clean up, the fun is done for now: `gcloud pubsub subscriptions delete test-sub && gcloud pubsub topics delete test`