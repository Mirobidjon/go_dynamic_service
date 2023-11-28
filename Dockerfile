FROM golang:1.21 as builder

RUN mkdir -p $GOPATH/src/gitlab-dev.soliqservis.uz/soliqservice/ss_go_dynamic_service
WORKDIR $GOPATH/src/gitlab-dev.soliqservis.uz/soliqservice/ss_go_dynamic_service 
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
    make linter && \
    make build && \
    mv ./bin/ss_go_dynamic_service /

FROM alpine:3.14

COPY --from=builder ss_go_dynamic_service .

RUN apk update && apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Tashkent /etc/localtime

ENTRYPOINT ["/ss_go_dynamic_service"]

