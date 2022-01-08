# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS app-builder

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata


##
## Deploy
##
FROM scratch
COPY --from=app-builder /go/bin/mqtt2ws /mqtt2ws
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 1883
ENTRYPOINT ["/mqtt2ws"]

