FROM golang:1.25.1 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o librarian ./cmd/librarian/main.go

EXPOSE 50052

CMD ["./librarian"]