package routers

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
