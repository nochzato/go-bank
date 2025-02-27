# Build stage
FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

# Run stage
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .
COPY db/migration ./db/migration

EXPOSE 8080 9090

CMD [ "/app/main" ]
