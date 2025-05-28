FROM golang:1.21

# Install build dependencies
RUN apt-get update && apt-get install -y \
  gcc \
  libc6-dev \
  libsqlite3-dev \
  pkg-config

# Set working directory
WORKDIR /app

# Copy go.mod & go.sum first (agar cache tidak rusak)
COPY go.mod go.sum ./

# Enable CGO for go mod if needed
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

# Download dependencies
RUN go mod tidy

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o app .

# Run the app
CMD ["./app"]
