

# Use an official lightweight Go image
FROM golang:1.24.1-alpine

# Set working directory
WORKDIR /app


# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./
RUN go mod download


# Copy the source code
COPY . .


# Build the Go binary
RUN go build -o home-app ./go_api_server
RUN go build -o home-app ./go_logging_server

# Run the app
CMD ["/app/home-app"]