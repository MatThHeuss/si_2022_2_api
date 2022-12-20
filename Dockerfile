FROM golang:1.18-alpine

LABEL maintainer="matheus alencar"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

ENTRYPOINT [ "tail", "-f", "/dev/null" ]