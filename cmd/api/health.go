package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthCheckHandler(c *gin.Context) {
	data := gin.H{
		"status": "ok",
		"env":    "development",
	}

	c.JSON(http.StatusOK, data)
}
