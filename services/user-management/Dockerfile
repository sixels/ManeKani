FROM golang:1.22-alpine as base

WORKDIR /usr/src/manekani-ums

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]