FROM golang:1.13-alpine3.10 AS build-env

RUN apk update && apk upgrade && \
   apk add --no-cache bash git pkgconfig gcc g++ libc-dev

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org

ENV TZ Europe/Amsterdam

WORKDIR /go/src/app

ADD . /go/src/app

# Because of how the layer caching system works in Docker, the go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

# set crosscompiling fla 0/1 => disabled/enabled
ENV CGO_ENABLED=0
# compile linux only
ENV GOOS=linux
# run tests
RUN go test ./... -covermode=atomic

RUN go build -v -ldflags '-w -s -linkmode external -extldflags -static' -a -installsuffix cgo -o /request-debug-endpoint .

FROM scratch

# important for time conversion
ENV TZ Europe/Amsterdam

WORKDIR /
ENV PATH=/
ENV SERVICECONFIG=/

# Import from builder.
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /request-debug-endpoint /