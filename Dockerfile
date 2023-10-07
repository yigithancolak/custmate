# Build stage
FROM golang:1.20.7-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bin/custmate server.go

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/production.env .
COPY --from=builder /app/bin/custmate .

ENV APP_ENV=production


EXPOSE 8080
CMD [ "/app/custmate" ]