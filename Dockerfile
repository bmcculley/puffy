FROM golang:1.16.5

MAINTAINER Michelle Noorali "michelle@deis.com"

COPY . /go/src/github.com/bmcculley/puffy

RUN go install github.com/bmcculley/puffy@master

ENTRYPOINT ["/go/bin/puffy"]

