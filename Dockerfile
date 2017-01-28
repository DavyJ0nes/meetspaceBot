FROM golang:1.7-onbuild
MAINTAINER DavyJ0nes <davy.jones@me.com>
ENV UPDATED_ON: 28-01-2017

RUN mkdir /app
COPY . /app/
VOLUME .
WORKDIR /app
CMD go test -v ./...
