# gcp-examples

In this repo I plan to collect some basic examples of how to use Go to interact with GCP.I plan to start with the
basics and eventually start composing some of the services together to build more "real" things.

Everything that is done in this repo is done so within the free-tier of GCP. Every example will include explicit
instructions on how to spin up and down the resources used.

## Assumptions

Here are the list of things you will need to do before you run any of the examples.

- You have a [Go version installed(1.11+)](https://golang.org/dl/) with
[modules turned on](https://github.com/golang/go/wiki/Modules#how-to-define-a-module).

- You have a GCP account. Great news, if you don't already have one it is free to sign up
[here](https://cloud.google.com/).

- You have a project created in GCP. One should be created for you by default. The import piece here is that your
project will have an id associated with it. For all of the examples provided you will need to have an environment
variable called `GCP_PROJECT_ID` set to the value of your id.

## Getting started

1. If you have not already, set an environment variable called `GCP_PROJECT_ID` to your GCP project id. You can see
this value in the GCP console.

2. Download `gcloud`. This is the command line tool used to work with GCP. I use ubuntu, so I just used the snap to
install it: `sudo snap install google-cloud-sdk --classic`

3. Login to your account: `gcloud auth login`

4. Allow local apps to auth easily with GCP: `gcloud auth application-default login`

5. Set what project you are working in. You can get this from your home dashboard: `gcloud config set project $GCP_PROJECT_ID`

6. Lastly, download the Go dependencies if you need them: `go mod download`

7. You are now ready to have some fun!


## Index

1. [Pub Sub](pubsub)
    - [Producer](pubsub/producer)
    - [Subscriber](pubsub/subscriber)

## Disclaimer

I am just a dude. I don't work with GCP on the regular and am learning as I go. Feel free to contribute if you see any
anti-patterns, or anything of the like.