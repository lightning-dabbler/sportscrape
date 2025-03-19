FROM golang:1.24.0-alpine3.21
RUN apk update && apk upgrade && apk add make && apk add chromium
RUN go install github.com/vektra/mockery/v2@v2.53.3
RUN export PATH=/usr/bin/chromium-browser:$PATH
