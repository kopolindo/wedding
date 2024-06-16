# Stage 1: Build the Go binary
FROM golang:1.22.4 as builder
WORKDIR /app
COPY backend .
RUN go mod download
RUN go build -o /server .

# Stage 2: Copy the binary to a smaller image
FROM alpine:latest
RUN apk update
RUN apk add --no-cache libc6-compat
RUN apk add --no-cache shadow

COPY .env /tmp/.env
RUN set -a
RUN . /tmp/.env
ARG GID
ARG UID
RUN addgroup -g $GID pi
RUN adduser -D -u $UID -G pi pi

RUN mkdir /var/log/backend
RUN chown -R pi:pi /var/log/backend
RUN chmod -R 755 /var/log/backend

USER pi
WORKDIR /app
COPY --from=builder /server /app/server
EXPOSE 8080
ENTRYPOINT [ "./server", "-debug" ]