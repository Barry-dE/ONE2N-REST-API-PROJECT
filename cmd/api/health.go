package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	statusHealthy   = "healthy"
	statusUnhealthy = "unhealthy"
	statusUp        = "UP"
	statusDown      = "DOWN"
)

type healthCheckResponse struct {
	Status         string `json:"status"`
	Env            string `json:"env"`
	DatabaseStatus string `json:"database_status"`
	TimeStamp      string `json:"timestamp"`
	// TO-DO: Redis status
}

func checkDatabaseStatus(db *sql.DB) (string, error) {
	if err := db.Ping(); err != nil {
		return statusDown, err
	}

	return statusUp, nil
}

func (app *application) healthCheckHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		dbStatus, err := checkDatabaseStatus(db)
		status := statusHealthy
		if err != nil {
			status = statusUnhealthy
		}

		healthData := healthCheckResponse{
			Status:         status,
			Env:            app.config.env,
			DatabaseStatus: dbStatus,
			TimeStamp:      time.Now().Format(time.RFC3339),
		}

		c.JSON(http.StatusOK, gin.H{
			"data": healthData,
		})
	}
}
