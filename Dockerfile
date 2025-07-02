FROM golang:1.24.2-alpine3.21
RUN apk update && apk upgrade && apk add make && apk add chromium
RUN go install github.com/vektra/mockery/v3@v3.5.0
RUN export PATH=/usr/bin/chromium-browser:$PATH
