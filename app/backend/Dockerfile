FROM golang:1.24

WORKDIR /go/src/app/backend

RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]