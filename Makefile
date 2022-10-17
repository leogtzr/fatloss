# fatloss Makefile is used to drive the build and installation of fatloss
# this is meant to be used with a local copy of code repository.

test:
	go test ./...

clean:
	@echo "Cleaning up build junk"

build:
	go build

cover:
	go test -v ./... -coverprofile ./cp.out
	go tool cover -html=cp.out
