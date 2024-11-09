# Start with the Go base image
FROM golang:1.19-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module and sum files first to leverage Docker layer caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o app .

# Expose the port that the application will run on
EXPOSE 8080

# Start the application
CMD ["./app"]
