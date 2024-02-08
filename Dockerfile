FROM golang:1.21 as golayer

ADD go.mod /go/src/github.com/Alevsk/rmm/go.mod
ADD go.sum /go/src/github.com/Alevsk/rmm/go.sum

WORKDIR /go/src/github.com/Alevsk/rmm/

RUN go mod download

ADD . /go/src/github.com/Alevsk/rmm/

ENV CGO_ENABLED=0

RUN go build -ldflags "-w -s" -a -o rmm ./cmd/rmm

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.7

MAINTAINER Lenin Alevski "lenin@alevsk.com"

WORKDIR /app

COPY --from=golayer /go/src/github.com/Alevsk/rmm/rmm /app/

ENTRYPOINT ["/app/rmm"]
