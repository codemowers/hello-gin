package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files (for CSS, JS, etc.)
	router.Static("/static", "./static")

	metricRouter := gin.Default()
	metrics := ginmetrics.GetMonitor()
	metrics.SetMetricPath("/metrics")
	metrics.UseWithoutExposingEndpoint(router)
	metrics.Expose(metricRouter)

	router.GET("/", func(c *gin.Context) {
		// Get environment variables
		username := "suvakelbas"


		podName := os.Getenv("MY_POD_NAME")
		if podName == "" {
			podName = "Unknown Pod"
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": username,
			"podName":  podName,
		})
	})

	go func() {
		_ = metricRouter.Run(":8080")
	}()

	router.Run(":8000")
}
