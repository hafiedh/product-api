FROM golang:latest as builder
LABEL MAINTAINER="Muhammad Hafiedh"

WORKDIR /go/src/api-product
COPY . .

RUN go mod download && \
    go mod verify


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/api-product .


FROM alpine:latest
RUN apk update && \
    adduser -D appuser

COPY --from=builder /go/bin/api-product /app/api-product
COPY --from=builder /go/src/api-product/.env /app/.env

USER appuser

WORKDIR /app

EXPOSE 8090

ENTRYPOINT ["/app/api-product", "-env", "/app/.env"]