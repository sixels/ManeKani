FROM golang:1.19-alpine as dev


WORKDIR /app/src/manekani

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .


RUN go build -o ./manekani/bin .

EXPOSE 8081

CMD ["./manekani"]