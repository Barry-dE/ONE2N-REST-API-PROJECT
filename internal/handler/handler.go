package handler

import (
	"net/http"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store repository.Storage
}

func NewHandler(store repository.Storage) *Handler {
	return &Handler{
		store: store,
	}
}

// Represents HTTP payload for student creation
type CreateStudent struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,gte=0,lte=150"`
	Sex       string `json:"sex" binding:"required,oneof=M F Other"`
}

// CreateStudentHandler handles student creation
func (h *Handler) CreateStudentHandler(c *gin.Context) {

	var payload CreateStudent

	// validate request
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Convert the incoming payload to the domain model
	student := &repository.Student{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Age:       payload.Age,
		Sex:       payload.Sex,
	}

	// Persist student data in the database
	err = h.store.Students.Create(c.Request.Context(), student)
	if err != nil {
		switch err {
		case repository.ErrDuplicateEmail:
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		}

		return
	}

	c.JSON(http.StatusCreated, student)
}
