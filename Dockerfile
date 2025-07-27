FROM golang:1.23.1-alpine AS builder

COPY . /github.com/UraharaKiska/go-auth/source/
WORKDIR /github.com/UraharaKiska/go-auth/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/UraharaKiska/go-auth/source/bin/auth .

CMD ["./auth"]