FROM alpine:latest
RUN apk update
RUN apk add --no-cache libc6-compat
RUN apk add --no-cache shadow

COPY .env /tmp/.env
RUN set -a
RUN . /tmp/.env
ARG GID
ARG UID
RUN addgroup -g $GID alex
RUN adduser -D -u $UID -G alex alex

RUN mkdir /var/log/backend
RUN chown -R alex:alex /var/log/backend
RUN chmod -R 755 /var/log/backend

USER alex
WORKDIR /app
COPY ./releases/server /app/server
EXPOSE 8080
ENTRYPOINT [ "./server", "-debug" ]