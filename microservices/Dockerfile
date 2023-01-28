FROM golang:1.18-alpine as builder

RUN apk upgrade --update-cache --available
RUN apk add --no-cache \
        gcc \
        musl-dev 

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /app/monolithique /app/main.go

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/monolithique ./

ENTRYPOINT ["/app/monolithique"]