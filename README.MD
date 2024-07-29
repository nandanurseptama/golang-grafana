# Golang Demo App with Grafana

This is demo application contains example how to send log data to loki and prometheus

## How to run

- clone repo

```bash
git clone https://nandanurseptama/golang-grafana.git
```

- pull docker necessary docker images

```bash
docker pull grafana/grafana:latest
docker pull grafana/loki:latest
docker pull grafana/promtail:latest
docker pull prom/prometheus:latest
docker pull golang:1.22-alpine # image version depends on your go version
```

- create `volumes` folder inside project

- build demo app with docker file

```bash
docker build -t your-docker-image-name .
```

- run docker compose
```bash
docker compose up -d
```

- access grafana and add loki and prometheus datasource
- import dashboard from exporter folder