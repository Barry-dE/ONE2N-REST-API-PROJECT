package handler

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/middleware"
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

type CreateStudentpayload struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,gte=0,lte=150"`
	Sex       string `json:"sex" binding:"required,oneof=M F Other"`
}

type updateStudentPayload struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email" binding:"omitempty,email"`
	Age       *int    `json:"age" binding:"omitempty,gte=0,lte=150"`
	Sex       *string `json:"sex" binding:"omitempty,oneof=M F Other"`
}

func (h *Handler) CreateStudentHandler(c *gin.Context) {

	var payload CreateStudentpayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		// To-do: standardize all json responses
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	student := &repository.Student{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Age:       payload.Age,
		Sex:       payload.Sex,
	}

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

func (h *Handler) GetStudentByID(c *gin.Context) {

	student := middleware.GetPostFromCtx(c)
	if student == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Student not found in context"})
		return
	}
	c.JSON(http.StatusOK, student)

}

func (h *Handler) UpdateStudentHandler(c *gin.Context) {

	student := middleware.GetPostFromCtx(c)
	if student == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Student not found in context"})
		return
	}

	var payload updateStudentPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentValue := reflect.ValueOf(student).Elem()
	payloadValue := reflect.ValueOf(payload).Elem()

	updated := false

	for i := 0; i < payloadValue.NumField(); i++ {
		field := payloadValue.Field(i)

		if !field.IsNil() {
			studentField := studentValue.Field(i)
			studentField.Set(field.Elem())
			updated = true
		}
	}

	if !updated {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No update fields provided"})
		return
	}

	err = h.store.Students.Update(c.Request.Context(), student)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		}
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *Handler) DeleteStudentHandler(c *gin.Context) {

}
