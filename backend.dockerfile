# Stage 1: Build the Go binary
FROM golang:latest as builder
WORKDIR /app
COPY backend .
RUN go mod download
RUN go build -o server .

# Stage 2: Copy the binary to a smaller image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]