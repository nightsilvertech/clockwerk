# Build clockwerk scheduler engine
FROM golang:alpine AS builder

RUN apk --no-cache add git mercurial
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go clean --modcache
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o clockwerk .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/clockwerk .

# Build a small image
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /dist/clockwerk /

# Command to run
ENTRYPOINT ["/clockwerk"]
