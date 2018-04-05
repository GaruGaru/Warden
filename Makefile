.PHONY: all clean

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

AUTHOR=garugaru
BINARY=warden
BUILD=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.Build=$(BUILD)"

build:
	go build -o $(BINARY) $(LDFLAGS)

clean:
	-rm $(BINARY)

docker_build:
	docker build -t "$(AUTHOR)/$(BINARY):$(BUILD)"
