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
RUN git config --global url."https://gitlab-ci-token:Qze6NiCgRcLJyyrhUNxn@gitlab.com/".insteadOf "https://gitlab.com/"
RUN go env -w GOPRIVATE=gitlab.com/nbdgocean6
RUN go clean --modcache
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o nobita-promo-scheduler .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/nobita-promo-scheduler .

# Build a small image
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /dist/nobita-promo-scheduler /
ENV TZ=Asia/Jakarta

# Command to run
ENTRYPOINT ["/nobita-promo-scheduler"]
