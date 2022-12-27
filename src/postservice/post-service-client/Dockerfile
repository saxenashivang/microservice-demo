FROM golang:alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/post-service-client

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN make tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/post-service-client/post-service-client /post-service-client
ENTRYPOINT ["/post-service-client"]
CMD []
