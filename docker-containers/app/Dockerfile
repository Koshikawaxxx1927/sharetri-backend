FROM golang:1.21

WORKDIR /app

COPY src/go.mod .
COPY src/go.sum .
COPY src/main.go .

RUN go mod tidy

CMD [ "go", "run", "main.go" ]