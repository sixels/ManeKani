FROM golang:1.19-alpine

VOLUME ["/data"]

WORKDIR /app/manekani

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /manekani-server

EXPOSE 8081

CMD ["/manekani-server"]