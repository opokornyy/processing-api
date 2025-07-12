FROM golang:1.24 AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . . 

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/app

FROM scratch

COPY --from=builder /app/main /main

EXPOSE 8080

# Run
CMD ["/main"]