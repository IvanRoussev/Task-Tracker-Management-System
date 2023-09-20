FROM golang:1.19
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
EXPOSE 9010

CMD ["/task-manager"]