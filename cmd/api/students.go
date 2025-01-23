package main

import (
	"net/http"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/store"
	"github.com/gin-gonic/gin"
)

// CreateStudentRequest represents the payload for student creation
type CreateStudent struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,gte=0,lte=150"`
	Sex       string `json:"sex" binding:"required,oneof=male female other"`
}

// CreateStudentHandler handles student creation
func (app *application) createStudentHandler(c *gin.Context) {

	var payload CreateStudent

	// validate request
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	// Convert the incoming payload to the domain model (Student)
	student := &store.Student{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Age:       payload.Age,
		Sex:       payload.Sex,
	}

	// Persist student data in the database
	err = app.store.Students.Create(c.Request.Context(), student)
	if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			c.JSON(http.StatusConflict, gin.H{"error": "email already exists"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		}

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         student.ID,
		"first_name": student.FirstName,
		"last_name":  student.LastName,
		"email":      student.Email,
		"age":        student.Age,
		"sex":        student.Sex,
		"created_at": student.CreatedAt,
		"updated_at": student.UpdatedAt,
	})
}

// curl -X POST http://localhost:8080/students \
//   -H "Content-Type: application/json" \
//   -d '{
//     "first_name": "John",
//     "last_name": "Doe",
//     "email": "john@example.com",
//     "age": 20,
//     "sex": "M"
