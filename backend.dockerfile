FROM alpine:latest
RUN apk update
RUN apk add --no-cache libc6-compat
RUN apk add --no-cache shadow

COPY .env /tmp/.env
RUN set -a && . /tmp/.env && set +a

ARG GID
ARG UID

RUN addgroup -g ${GID} alex && \
    adduser -D -u ${UID} -G alex alex
USER alex

WORKDIR /app
COPY ./releases/server /app/server
EXPOSE 8080
ENTRYPOINT [ "./server", "-debug" ]