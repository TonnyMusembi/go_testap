# Use Go 1.23 image to match your local version
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (to leverage Docker cache)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your application code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the built executable
CMD ["./main"]
