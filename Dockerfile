# build stage
FROM golang:1.9.1 AS builder
ADD . /src
WORKDIR /src
RUN go get -d -v ./... && \
    go build -ldflags "-linkmode external -extldflags -static" -o proxy

# final stage
FROM scratch
WORKDIR /app
COPY --from=builder /src/proxy /src/rules.json /app/

ENTRYPOINT [ "./proxy" ]
CMD ["./proxy"]
