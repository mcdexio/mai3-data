FROM golang:1.16.7-alpine3.14 as builder
RUN apk add build-base
WORKDIR /app
COPY . .
RUN go build -o mai3data-api main.go

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/mai3data-api .
ENTRYPOINT ["./mai3data-api"]