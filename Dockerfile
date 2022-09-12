FROM alpine:3.14.1

RUN apk update
RUN apk add tar unzip gstreamer