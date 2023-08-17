# Use the official Go image as the base image
FROM golang:1.21 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd

# Create a minimal final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous stage
COPY --from=build /app/app .

# Copy the configuration file into the container
COPY settings.yml .

# Command to run the executable
CMD ["./app"]
