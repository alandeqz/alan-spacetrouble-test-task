FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./src

FROM alpine AS alpine

RUN apk update && \
    apk upgrade && \
    apk add git

COPY --from=builder /app/main /app/main

WORKDIR /app

EXPOSE 8080

CMD ["./main"]
