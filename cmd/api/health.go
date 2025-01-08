package main

import (
	"net/http"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/env"
	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	data := gin.H{
		"status": "ok",
		"env":    env.GetString("ENV", "Development"),
	}

	c.JSON(http.StatusOK, data)
}
