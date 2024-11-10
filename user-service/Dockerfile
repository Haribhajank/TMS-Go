# Use an official Go image as the base
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o user-service

# Expose the port the service will run on
EXPOSE 8080

# Command to run the executable
CMD ["./user-service"]
