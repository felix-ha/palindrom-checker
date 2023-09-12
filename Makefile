.PHONY: test
test:
	go test -v ./...

.PHONY: build
build:
	go build -o palindromcheck main.go

.PHONY: docker_build
docker_build:
	docker build . -t felixhau/palindrom-checker:local

.PHONY: docker_run
docker_run:
	docker run --env STRING=rentner felixhau/palindrom-checker:local
	