FROM golang:1-alpine AS builder

MAINTAINER orginux
WORKDIR $GOPATH/src/container-hostname/

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/container-hostname

FROM scratch
COPY --from=builder /go/bin/container-hostname /go/bin/container-hostname
EXPOSE 8080
ENTRYPOINT ["/go/bin/container-hostname"]
