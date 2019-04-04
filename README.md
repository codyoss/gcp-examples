# gcloud pubsub

Produce and consume via `pubsub` in 10 easy to follow steps!

Example of how to produce messages to a `pubsub` topic from Go and consume them from your console with `gcloud`.
Everything that is done here is well within the free tier for `pubsub`. Directions include setup and teardown of all
GCP resources used in the project.

## Assumptions

- You have Go installed with modules turned on.(The future is now, embrace it!)
- You have a GCP Account created
- You have a project created in GCP

## Directions

1. Download the Go deps if you need them: `go mod download`

2. Install gcloud if you have not already. I use ubuntu so I just used the snap: `sudo snap install google-cloud-sdk --classic`

3. Login to your account: `gcloud auth login`

4. Allow local apps to auth easily with GCP: `gcloud auth application-default login`

5. Set what project you are working in. You can get this from your home dashboard: `gcloud config set project PROJECT_ID`

6. Create a pubsub topic. This is what you will send and recieve data from: `gcloud pubsub topics create test`

7. Create a pubsub subscriber. This is what you will use to pull message from the command line: `gcloud pubsub subscriptions create test-sub --topic test`

8. Run the provided code. You will need to swap in your `PROJECT_ID` again: `PROJECT_ID=some-project-id go run main.go`

9. Pull the data your app just produced! : `gcloud pubsub subscriptions pull test-sub`

10. Clean up, the fun is done for now: `gcloud pubsub subscriptions delete test-sub && gcloud pubsub topics delete test`