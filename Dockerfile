# Start from the 1.14 golang base image
FROM golang:1.14-alpine3.12

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go modm sum and src directory
COPY go.mod go.sum src ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy all files to docker
COPY . .

# Build the Go app
RUN go build -o main ./src/net/elau/gorestapi

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./src/net/elau/gorestapi/main"]