package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/handler"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/middleware"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/repository"

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
	config  Config
	handler *handler.Handler
}

func (app *application) mount(repo *repository.Storage) *gin.Engine {
	router := gin.Default()

	//custom middleware
	// router.Use(middleware.DisallowUnknownFieldsMiddleware())

	v1 := router.Group("/api/v1")
	{

		v1.GET("/healthCheck", app.healthCheckHandler)

		students := v1.Group("/students")
		{
			students.POST("/", app.handler.CreateStudentHandler)

			studentID := students.Group("/:studentID")
			studentID.Use(middleware.StudentContextMiddleware(repo))
			{
				studentID.GET("", app.handler.GetStudentByID)
				studentID.PATCH("", app.handler.UpdateStudentHandler)
			}
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
