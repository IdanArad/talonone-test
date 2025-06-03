FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the binary from /app, assuming main.go is at /app/main/main.go
RUN go build -o server ./main

EXPOSE 8080

CMD ["./server"]
