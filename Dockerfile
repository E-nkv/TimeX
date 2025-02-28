# Use the official Go image as a base
FROM golang:1.24.0

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application listens on
EXPOSE 8080

# Run the command to start your application when the container launches
CMD ["./main"]
