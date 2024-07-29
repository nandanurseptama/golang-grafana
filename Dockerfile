FROM golang:1.22-alpine

ENV PORT=8080
ENV PROMETHEUS_SERVER_PORT=2221
ENV LOG_FILE_PATH=/var/log/app.log

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]