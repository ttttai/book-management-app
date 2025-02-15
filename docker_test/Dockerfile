FROM golang:1.24

WORKDIR /docker_test

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD ["bash"]