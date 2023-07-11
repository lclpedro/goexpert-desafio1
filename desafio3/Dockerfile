FROM golang:latest

RUN mkdir /app
ENV GOPATH=/app/cmd

ADD . /app

WORKDIR /app

RUN go build -o ./cmd/main ./cmd
ENTRYPOINT ["./cmd/main"]
EXPOSE 8080