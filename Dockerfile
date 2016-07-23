FROM golang:1.6.2-alpine

MAINTAINER Kyle Banks

# Make the source code path
RUN mkdir -p /go/src/github.com/KyleBanks/glock

# Add all source code
ADD . /go/src/github.com/KyleBanks/glock

# Run the Go installer
RUN go install -v github.com/KyleBanks/glock

# Run the glock server and expose the port
ENTRYPOINT /go/bin/glock
EXPOSE 7887