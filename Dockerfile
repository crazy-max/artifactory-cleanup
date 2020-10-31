FROM --platform=${BUILDPLATFORM:-linux/amd64} tonistiigi/xx:golang AS xgo
FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:1.15-alpine AS builder

ARG VERSION=dev

ENV CGO_ENABLED 0
ENV GO111MODULE on
COPY --from=xgo / /

ARG TARGETPLATFORM
RUN go env

RUN apk --update --no-cache add \
    build-base \
    gcc \
    git \
  && rm -rf /tmp/* /var/cache/apk/*

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./
RUN go build -ldflags "-w -s -X 'main.version=${VERSION}'" -v -o artifactory-cleanup cmd/main.go

FROM --platform=${TARGETPLATFORM:-linux/amd64} alpine:latest

LABEL maintainer="CrazyMax"

RUN apk --update --no-cache add \
    ca-certificates \
    libressl \
    tzdata \
  && rm -rf /tmp/* /var/cache/apk/*

COPY --from=builder /app/artifactory-cleanup /usr/local/bin/artifactory-cleanup
RUN artifactory-cleanup --version

ENTRYPOINT [ "artifactory-cleanup" ]
