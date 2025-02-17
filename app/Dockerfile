FROM golang:1.24

WORKDIR /go/src/app

RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]