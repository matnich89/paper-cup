FROM golang:1.19 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/api/*.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]
