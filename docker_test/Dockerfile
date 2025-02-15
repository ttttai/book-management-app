FROM golang:1.24

WORKDIR /docker_test

RUN go install github.com/air-verse/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]