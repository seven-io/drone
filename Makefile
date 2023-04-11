build:
	@echo "Building drone-seven..."
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-seven
	@docker build -t drone-seven .
	@echo "Finished building drone-seven\n"

push:
	@echo "Pushing seven-io/drone to docker..."
	@docker push seven-io/drone
	@echo "Successfully pushed seven-io/drone to docker\n"