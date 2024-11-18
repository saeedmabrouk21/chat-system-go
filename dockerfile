# Use the official Golang image as the base image
FROM golang:1.20-alpine

RUN apk update && apk add bash

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .


# Copy the wait-for-it.sh script
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Build the Go app
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the Go app
CMD ["/bin/sh", "-c", "/wait-for-it.sh mysql:3306 -- ./main"]
