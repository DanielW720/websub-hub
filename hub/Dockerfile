FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o main ./src

EXPOSE 8080

CMD ["./main"]
