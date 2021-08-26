FROM golang:1.17.0-alpine3.14 AS base

ENV GOCACHE=/go/.go/cache GOPATH=/go/.go/path TZ=Europe/London

RUN apk update && apk upgrade && \
    apk add git && \
    apk add build-base

RUN GOBIN=/bin go get github.com/cespare/reflex

RUN ln -s /go /dateguess-api
WORKDIR /dateguess-api

# The build stage builds the binary from source
FROM base AS build

COPY go.mod go.sum ./

RUN go mod download

COPY cmd                            ./cmd
COPY internal                       ./internal
COPY pkg                            ./pkg

WORKDIR /dateguess-api/cmd/dateguess-api

RUN CGO_ENABLED=0 go build -o dateguess-api -ldflags -s

RUN apk add ca-certificates
RUN update-ca-certificates

FROM base AS debug

RUN go get github.com/go-delve/delve/cmd/dlv \
 && cp /go/.go/path/bin/dlv /usr/local/bin

WORKDIR /dateguess-api

EXPOSE 40000

ENV headless=true

CMD dlv debug src/*.go --listen=:40000 --headless=${headless} --api-version=2

# Main build
FROM scratch AS prod

WORKDIR /dateguess-api

COPY --from=build /dateguess-api/cmd/dateguess-api/dateguess-api       ./
COPY --from=build /etc/ssl/certs       /etc/ssl/certs

CMD ["/dateguess-api/dateguess-api"]
