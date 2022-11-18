FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/server/main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main" ]