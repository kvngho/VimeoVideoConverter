# Stage 1: Build the Go application
FROM golang:1.22-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go


# Stage 2: Create a minimal image with the binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .
COPY --from=build /app/config.yaml .

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the executable
CMD ["./main"]
