# Build Step
FROM golang:1.10 AS build

# Prerequisites and vendoring
RUN mkdir -p $GOPATH/src/github.com/GaruGaru/Warden
ADD . $GOPATH/src/github.com/GaruGaru/Warden
WORKDIR $GOPATH/src/github.com/GaruGaru/Warden
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o warden
RUN cp warden /

# Final Step
FROM alpine

# Base packages
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

# Copy binary from build step
COPY --from=build /warden /home/

# Define the ENTRYPOINT
WORKDIR /home
ENTRYPOINT ./warden
