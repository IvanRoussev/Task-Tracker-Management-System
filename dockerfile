FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download github.com/stretchr/testify@v1.7.0

RUN go mod download

COPY ../taskManager ./

RUN go build -o /app/taskManager

EXPOSE 8080

CMD [ "/app/taskManager" ]