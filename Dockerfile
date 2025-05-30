FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./

# Only include go.sum if it exists
# COPY go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]

