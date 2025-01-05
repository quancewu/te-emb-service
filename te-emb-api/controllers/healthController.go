package controllers

import (
	"context"
	"os"
	"runtime"
	"te-emb-api/initalizers"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

type HealthStatus struct {
	Status     string `json:"status"`
	Timestamp  string `json:"timestamp"`
	Version    string `json:"version"`
	Uptime     string `json:"uptime"`
	SystemInfo gin.H  `json:"system_info"`
	Services   gin.H  `json:"services"`
}

func HealthCheck(c *gin.Context) {
	// check Redis connection
	redisStatus := "ok"
	if err := initalizers.Redis.Ping(context.Background()).Err(); err != nil {
		redisStatus = "error"
	}

	health := HealthStatus{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   os.Getenv("APP_VERSION"), // get from env
		Uptime:    time.Since(startTime).String(),
		SystemInfo: gin.H{
			"go_version": runtime.Version(),
			"goroutines": runtime.NumGoroutine(),
			"os":         runtime.GOOS,
			"arch":       runtime.GOARCH,
		},
		Services: gin.H{
			"redis": redisStatus,
		},
	}

	c.JSON(200, health)
}
