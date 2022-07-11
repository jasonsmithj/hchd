FROM alpine:3.16.0

MAINTAINER jasonsmithj
ARG CONTAINER_PORT

RUN apk update \
 && apk upgrade \
 && apk add --update-cache --no-cache --virtual .create-user shadow \
 && groupadd -g 1000 jason \
 && useradd -u 1000 -g 1000 jason \
 && apk del .create-user \
 && apk add ca-certificates \
 && apk --update add tzdata \
 && rm -rf /var/cache/apk/*

USER jason
EXPOSE $CONTAINER_PORT
WORKDIR /app/jason/
COPY ./.artifacts/hchd-linux-amd64 ./app
