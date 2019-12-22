# We specify the base image we need for our
# go application
FROM golang:latest 
# We create an /go/src/berlin directory within our
# image that will hold our application source
# files
RUN mkdir /go/src/berlin
# We copy everything in the root directory
# into our /go/src/berlin directory
RUN go get -u github.com/golang/dep/cmd/dep
ADD . /go/src/berlin
COPY ./Gopkg.toml /go/src/berlin
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /go/src/berlin
# Add this go mod download command to pull in any dependencies
RUN dep ensure
# Installing redis server inside docker
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install redis-server -y
# we run go build to compile the binary
# executable of our Go program
RUN go build -o main .
# Give access to execute go main
RUN chmod +x /go/src/berlin/*
# Our start command which kicks off
# our newly created binary executable
ENTRYPOINT redis-server --daemonize yes && /go/src/berlin/main