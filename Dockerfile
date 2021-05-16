FROM golang:1.16.4-alpine3.13 AS builder
LABEL maintainer="orginux"
WORKDIR $GOPATH/src/echopod/
COPY main.go .
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux \
        go build -o /go/bin/echopod

FROM scratch
COPY --from=builder /go/bin/echopod /go/bin/echopod
EXPOSE 8080
ENTRYPOINT ["/go/bin/echopod"]
