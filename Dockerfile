FROM golang:1.19-alpine
LABEL authors="kiril"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go

CMD ["/app/main"]