# Use the official Golang image as the base image
FROM golang:1.21.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the project
RUN go build -o main ./cmd

# Expose the port the application runs on
EXPOSE 3000

# Command to run the executable
CMD ["./main"]