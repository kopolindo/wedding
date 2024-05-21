# Stage 1: Build the Go binary
FROM golang:latest as builder
WORKDIR /app
COPY backend .
RUN go mod download
RUN go build -o /server .

# Stage 2: Copy the binary to a smaller image
FROM alpine:latest
RUN apk add libc6-compat
RUN addgroup -S service && adduser -S service -G service
RUN apk update
RUN mkdir /var/log/backend
RUN chown -R service:service /var/log/backend
RUN chmod -R 755 /var/log/backend
USER service
WORKDIR /app
COPY --from=builder /server /app/server
EXPOSE 8080
USER service:service
ENTRYPOINT [ "./server", "-debug" ]