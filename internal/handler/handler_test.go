package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/handler"
	"github.com/Barry-dE/ONE2N-REST-API-PROJECT/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStudentStore is a mock implementation of the StudentRepository
type MockStudentStore struct {
	mock.Mock
}

func (m *MockStudentStore) Create(ctx context.Context, student *repository.Student) error {
	args := m.Called(ctx, student)
	return args.Error(0)
}

func TestCreateStudentHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		mockError      error
		expectedStatus int
	}{
		{
			name: "Successful student creation",
			requestBody: map[string]interface{}{
				"firstName": "John",
				"lastName":  "Doe",
				"email":     "john@example.com",
				"age":       20,
				"sex":       "M",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Validation error - missing required fields",
			requestBody: map[string]interface{}{
				"firstName": "John",
				"email":     "john@example.com",
				"age":       20,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Duplicate email error",
			requestBody: map[string]interface{}{
				"firstName": "Jane",
				"lastName":  "Doe",
				"email":     "jane@example.com",
				"age":       22,
				"sex":       "F",
			},
			mockError:      repository.ErrDuplicateEmail,
			expectedStatus: http.StatusConflict,
		},
		{
			name: "Internal server error",
			requestBody: map[string]interface{}{
				"firstName": "Jake",
				"lastName":  "Smith",
				"email":     "jake@example.com",
				"age":       25,
				"sex":       "M",
			},
			mockError:      errors.New("database error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create mock store
			mockStore := new(MockStudentStore)

			// Create storage with mock store
			storage := repository.Storage{
				Students: mockStore,
			}

			// Convert request body to JSON
			body, _ := json.Marshal(tc.requestBody)

			// Create a new HTTP request
			req, _ := http.NewRequest(http.MethodPost, "/students", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			// Mock the database call if applicable
			if tc.expectedStatus == http.StatusCreated || tc.expectedStatus == http.StatusConflict || tc.expectedStatus == http.StatusInternalServerError {
				mockStore.On("Create", mock.Anything, mock.AnythingOfType("*repository.Student")).Return(tc.mockError)
			}

			// Initialize handler
			h := handler.NewHandler(storage)
			h.CreateStudentHandler(c)

			// Assert the status code
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Verify that our expectations were met
			mockStore.AssertExpectations(t)
		})
	}
}
