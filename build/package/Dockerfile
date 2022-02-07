FROM golang:1.17 as builder

WORKDIR /go/src/app
COPY . .

RUN go build -o /build/bin/api ./cmd/api

FROM centos:7

WORKDIR /app

COPY web /app/web
COPY --from=builder /build/bin/api /build/bin/api
COPY configs/config.yml /etc/api.yml

ENTRYPOINT ["/build/bin/api", "--configPath=/etc/api.yml"]