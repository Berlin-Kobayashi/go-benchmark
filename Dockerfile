FROM golang:1.10 as go
WORKDIR /go/src/github.com/DanShu93/go-benchmark
COPY . .

RUN CGO_ENABLED=0 go build server.go

FROM alpine

WORKDIR /app

COPY --from=go /go/src/github.com/DanShu93/go-benchmark/server .

ARG API

CMD /app/server --api $API
