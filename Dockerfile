FROM golang:1.21.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/go-docker-app"]
