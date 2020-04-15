FROM golang:1.14.2-alpine3.11

ENV SOURCE /go/src/github.com/jonnay101/icon/cmd/
ENV GO111MODULE=on

COPY . /icon

WORKDIR /icon/cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ENV PORT 8080

EXPOSE 8080

ENTRYPOINT ["/cmd/main"]