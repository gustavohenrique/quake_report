FROM golang:1.23-alpine

WORKDIR /app
COPY . .

CMD ["go", "run", "main.go"]

