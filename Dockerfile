FROM golang:1.21

# Install dependencies for CGO and SQLite
RUN apt-get update && apt-get install -y gcc libc6-dev libsqlite3-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app .

CMD ["./app"]
