FROM golang:1.21.5

WORKDIR /src/api

COPY . .

RUN go build ./cmd/main/

CMD [ "go", "run", "./cmd/main/main.go" ]

