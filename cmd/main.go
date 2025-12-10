package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	metricRouter := gin.Default()
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.Use(router)
	metrics.Expose(metricRouter)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello " + os.Getenv("USERNAME") + " from " + os.Getenv("MY_POD_NAME"))
	})

	go func() {
		_ = metricRouter.Run(":8080")
	}()

	router.Run(":8000")
}
