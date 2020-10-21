FROM golang:1-alpine AS builder

MAINTAINER orginux
WORKDIR $GOPATH/src/container-info/

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/container-info

FROM scratch
COPY --from=builder /go/bin/container-info /go/bin/container-info
EXPOSE $PORT
ENTRYPOINT ["/go/bin/container-info"]
