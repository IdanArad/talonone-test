FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./main ./main

WORKDIR /app/main

RUN go build -o ../server .

EXPOSE 8080

CMD ["../server"]

