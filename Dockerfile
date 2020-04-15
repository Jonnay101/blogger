FROM golang:1.14.2-alpine3.11 as build-service

# Turn go mods on
ENV GO111MODULE=on

# Caching the go dependencies
WORKDIR /icon
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all source code
COPY . /icon

# Build the service
ENV PORT 8080
WORKDIR /icon/cmd/icon/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Build smaller binary
FROM scratch
COPY --from=build-service /icon/cmd/icon/icon /usr/bin/icon
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/usr/bin/icon"]