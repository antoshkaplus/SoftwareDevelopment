FROM ubuntu:latest

WORKDIR /usr/src/app

RUN apt update
RUN apt install -y libzmq3-dev libczmq-dev
RUN apt install -y golang-go
RUN apt install ca-certificates -y
RUN update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build
RUN mkdir ../tmp

CMD ["./fruit", "server"]
