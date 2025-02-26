FROM golang:1.24.0-alpine3.21
RUN apk update && apk upgrade && apk add make && apk add chromium
RUN export PATH=/usr/bin/chromium-browser:$PATH
