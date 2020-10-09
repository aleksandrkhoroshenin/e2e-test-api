FROM golang:1.14.7 as builder

RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY . .

RUN go mod download

RUN cd cmd/gateway && CGO_ENABLED=0 go build

RUN cd migrations && CGO_ENABLED=0 go build
