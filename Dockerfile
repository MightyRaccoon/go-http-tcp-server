# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR \app
COPY . .
RUN go mod download

RUN go build -o /go-http-tcp-server

EXPOSE 80

CMD ["/go-http-tcp-server"]
