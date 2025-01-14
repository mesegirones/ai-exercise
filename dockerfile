FROM golang:1.22 AS builder

# Move to working directory /app
WORKDIR /app

# Copy and download dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main ./app

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/main .

# Run the web service on container startup.
CMD ["/app/main"]