build:
	@echo "Building drone-sms77..."
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sms77
	@docker build -t drone-sms77 .
	@echo "Finished building drone-sms77\n"

push:
	@echo "Pushing sms77io/drone to docker..."
	@docker push sms77io/drone
	@echo "Successfully pushed sms77io/drone to docker\n"