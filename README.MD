# Prometheus Metric Collection

This is demo application to collect metrics from your golang service

## Initialization

In this article, we will use grafana and prometheus inside docker. So you need to pull grafana and prometheus images.

```bash
docker pull grafana/grafana:latest
docker pull prom/prometheus:latest
```

Create a simple api service, but in this article i will use `golang` and `gin` library (http server).

```bash
mkdir golang-grafana
cd golang-grafana
go mod init github.com/your-username/your-repo-name.git
```

## Coding

Create `main.go` file inside root project.

```go
// get env variable, if not found use `defaultValue`
func safeEnv(name string, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}
	return val
}

func main(){
    port := safeEnv("PORT", "8080") // service port
	promServerPort := safeEnv("PROMETHEUS_SERVER_PORT", "2221") // prometheus client port

    // app http server
    r := gin.Default()
    // prometheus client server
    promServer := gin.Default()
}
```
