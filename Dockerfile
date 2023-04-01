# build golang executable
FROM golang:1.20-alpine AS builder
WORKDIR /go/src/github.com/ahdark-services/ahdark-me-redirector

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/ahdark-me-redirector .

# Build a small image
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Update ca-certificates
RUN update-ca-certificates

# Copy our static executable
COPY --from=builder /go/bin/ahdark-me-redirector /go/bin/ahdark-me-redirector

# Run the ahdark-me-redirector binary.
ENTRYPOINT ["/go/bin/ahdark-me-redirector"]

# Expose port 8080 to the outside world
EXPOSE 8080
