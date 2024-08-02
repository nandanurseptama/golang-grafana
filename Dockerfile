FROM golang:1.22-alpine

ENV PORT=8080

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]