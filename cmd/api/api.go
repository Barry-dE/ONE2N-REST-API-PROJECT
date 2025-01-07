package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	addr string
	env  string
}

type application struct {
	config Config
}

func (app *application) mount() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{

		v1.GET("/healthCheck", app.healthCheckHandler)

		v1.Group("/students")
		{

		}
	}

	return router
}

func (app *application) run(handler *gin.Engine) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return srv.ListenAndServe()

}
