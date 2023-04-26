
FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main main.go

EXPOSE 8000

CMD ["/app/main"]
