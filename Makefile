# To get/read all environment variable
# from .env file
include .env
export

# To open the webhook to the world
online:
	ssh -R 80:localhost:9091 localhost.run

# To run the Webhook watcher
run:
	go run main.go

