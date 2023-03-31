FROM golang:1.20-alpine


WORKDIR /usr/src/manekani

RUN go install github.com/cosmtrek/air@latest

COPY ../go.mod ../go.sum ./
RUN go mod download && go mod verify

EXPOSE 8081

CMD ["air"]