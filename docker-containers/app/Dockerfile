FROM golang:1.21

WORKDIR /app

COPY ../../go.mod .
COPY ../../go.sum .
COPY ../../main.go .

RUN go mod tidy

CMD [ "go", "run", "main.go" ]