FROM golang:1.21-alpine3.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main

FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]