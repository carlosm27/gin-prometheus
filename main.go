package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	prometheus.MustRegister(pingCounter)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/ping", ping)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run()
}

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(c *gin.Context) {
	pingCounter.Inc()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
