FROM golang:1.20-alpine as base
WORKDIR /app

ENV HOST=0.0.0.0
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Await dependencies script
RUN apk add --no-cache bash
RUN wget https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh -O /usr/local/bin/wait-for-it && chmod a+x  /usr/local/bin/wait-for-it

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
ENV GOPRIVATE=github.com/pathai95441/*

COPY .netrc /root/.netrc
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o go-boilerplate -a ./cmd/go-boilerplate

### Production
FROM alpine:latest

ENV HOST=0.0.0.0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/go-boilerplate /usr/local/bin/go-boilerplate
COPY --from=builder /app/configs /usr/local/bin/configs
COPY --from=builder /app/docs /usr/local/bin/docs

WORKDIR /usr/local/bin/

EXPOSE 8080

CMD ["./go-boilerplate"]
