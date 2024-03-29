FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64

# Move to working directory /build
WORKDIR /app
# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
ADD . .

# Build the application
RUN go build -o app .


# Copy binary from build to main folder
ENTRYPOINT ["./app"]

# Port litsen
EXPOSE 8080