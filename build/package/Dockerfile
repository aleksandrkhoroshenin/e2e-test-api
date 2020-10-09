FROM golang:1.14.7 as builder

RUN mkdir -p /opt/app
WORKDIR /opt/app

COPY . .

# Install golangci-lint
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

RUN go mod download

RUN cd cmd/gateway && CGO_ENABLED=0 go build

RUN cd migrations && CGO_ENABLED=0 go build

RUN ./bin/golangci-lint run

