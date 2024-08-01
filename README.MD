# Metric Collection with Prometheus

## Initialization

In this article i will show you, how to collect metric inside your go service with Prometheus.
I've created demo api with Golang, so you can clone it.

```bash
git clone git@github.com:nandanurseptama/golang-grafana.git
git checkout demo-app # checkout to demo-app branch
cd golang-grafana
```

First, we need to install prometheus library in go mod

```bash
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
```

Open `main.go` and add new route `/metrics`, this route will used by prometheus to show your api metrics.

```go
r.POST("/api/login", routers.Login(logClient))
r.POST("/api/register", routers.Register(logClient))
// this is will be used by prometheus
r.GET("/metrics", func(ctx *gin.Context) {
    promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
})
```

then run `main.go` file

```bash
go run main.go
```

You can access it `http://localhost:8080/metrics` to see metric datas.

## Create custom metric

In this article, we will custom metrics.

So first, you can create new file `route_metric.go` inside `routers` folder and paste code below.

```go
import (
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
// total user registered counter
totalSuccessRegisterCounter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "total_success_register",
    Help: "Number of successful registrations",
})

// total failed registered counter
totalFailedRegisterCounter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "total_failed_register",
    Help: "Number of failed registrations",
})

// total success login counter
totalSuccessLoginCounter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "total_success_login",
    Help: "Number of successful login",
})

// total failed login counter
totalFailedLoginCounter = promauto.NewCounter(prometheus.CounterOpts{
    Name: "total_failed_login",
    Help: "Number of failed login",
})
)
```

At this point, when you run `main.go` and access `/metrics` path you can see new metrics name that we added above. Now, open file `routers/login.go` and call counter to increment value depending on success or failed case

1. When parsing request body failed, increment `totalFailedLoginCounter`

```go
// bind request body
if err := ctx.ShouldBindJSON(&body); err != nil {
    ctx.JSON(
        http.StatusInternalServerError,
        map[string]any{"message": "failed to bind request body"},
    )
    ctx.Abort()
    totalFailedLoginCounter.Inc() // Increment counter
    return
}
```

2. When email or password empty, increment `totalFailedLoginCounter`

```go
// guard when email and password empty
// send response error
if body.Email == "" || body.Password == "" {
    ctx.JSON(
        http.StatusBadRequest,
        map[string]any{"message": "email and password required"},
    )
    ctx.Abort()
    totalFailedLoginCounter.Inc()
    return
}
```

3. When email not registered, increment `totalFailedLoginCounter`

```go
// guard when user not registered
if findUser.Email == "" {
    ctx.JSON(
        http.StatusBadRequest,
        map[string]any{"message": "email or password not match"},
    )
    ctx.Abort()
    totalFailedLoginCounter.Inc()
    return
}
```

4. When password not match, increment `totalFailedLoginCounter`

```go
// guard when password not match
if findUser.Password != body.Password {
    ctx.JSON(
        http.StatusBadRequest,
        map[string]any{"message": "email or password not match"},
    )
    ctx.Abort()
    totalFailedLoginCounter.Inc()
    return
}
```

5. When email and password match, increment `totalSuccessLoginCounter`

```go
ctx.JSON(http.StatusOK, map[string]any{"message": "OK"})
ctx.Abort()
totalSuccessLoginCounter.Inc()
```

After that, you can run `main.go` and try to hit `/api/login` path with success or failed case. You can see the `total_success_login` or `total_failed_login` value will change.


## Visualization Prometheus Metrics with Grafana