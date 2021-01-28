FROM golang:1.15

WORKDIR /app
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY main.go /app/main.go
COPY statik/statik.go /app/statik/statik.go
RUN go build -o app

CMD ["./app"]
