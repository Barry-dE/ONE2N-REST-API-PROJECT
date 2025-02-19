package middleware

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/repository"
	"github.com/gin-gonic/gin"
)

const StudentContextKey string = "student"

func StudentContextMiddleware(store *repository.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("studentID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			c.Abort()
			return
		}

		student, err := store.Students.GetByID(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			c.Abort()
			return
		}

		c.Set(StudentContextKey, student)

		c.Next()
	}
}

func GetPostFromCtx(c *gin.Context) *repository.Student {
	value, exists := c.Get(StudentContextKey)
	if !exists {
		return nil
	}

	student, ok := value.(*repository.Student)
	if !ok {
		return nil
	}

	return student
}
