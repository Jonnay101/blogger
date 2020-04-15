FROM golang:1.14.2-alpine3.11

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
WORKDIR /icon/cmd/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Expose the port and service entrypoint
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/cmd/main"]