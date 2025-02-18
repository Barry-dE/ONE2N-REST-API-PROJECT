package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/middleware"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/store"
	"github.com/gin-gonic/gin"
)

type dbConfig struct {
	addr               string
	maxOpenConnections int
	maxIdleConnections int
	MaxLifetime        time.Duration
	maxIdleTime        time.Duration
}

type Config struct {
	addr string
	env  string
	db   dbConfig
}

type application struct {
	config Config
	store  store.Storage
}

func (app *application) mount() *gin.Engine {
	router := gin.Default()

	//custom middleware

	router.Use(middleware.DisallowUnknownFields())

	v1 := router.Group("/api/v1")
	{

		v1.GET("/healthCheck", app.healthCheckHandler)

		students := v1.Group("/students")
		{
			students.POST("/", app.createStudentHandler)
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

	log.Printf(" Server running on %s", app.config.addr)

	return srv.ListenAndServe()

}
