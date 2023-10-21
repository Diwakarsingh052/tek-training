package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"service-app/middlewares"
	"service-app/models"
	"service-app/models/mockmodels"
	"testing"
)

func TestSignup(t *testing.T) {
	// NewUser struct initialization
	nu := models.NewUser{
		Name:     "Diwakar",
		Email:    "diwakar@email.com",
		Password: "password",
	}
	mockUser := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:         "Diwakar",
		Email:        "diwakar@email.com",
		PasswordHash: "jaldlajjasdf",
	}

	tt := [...]struct {
		name             string // Name of the test case
		body             any    // Body to send to request
		expectedStatus   int    // Expected status of the response
		expectedResponse string // Expected response body
		expectedUser     models.User
		mockUserService  func(m *mockmodels.MockService)
	}{
		{
			name:             "OK",
			body:             nu,
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Diwakar","email":"diwakar@email.com"}`,
			expectedUser:     mockUser,
			//set expectations inside it
			mockUserService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Eq(nu)).
					Times(1).Return(mockUser, nil)

			},
		},
		{
			name: "Fail_NoEmail",
			body: models.NewUser{
				Name:     "testuser",
				Password: "password",
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"msg":"please provide Name, Email and Password"}`,
			mockUserService: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)
			//this func give us the mock implementation of the interface
			mockService := mockmodels.NewMockService(ctrl)
			s := models.NewStore(mockService)

			// Apply the mock to the user service.
			tc.mockUserService(mockService)
			// Create a new Gin router.
			router := gin.New()
			h := handler{
				s: s,
			}
			ctx := context.Background()
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Register an endpoint and its handler with the router.
			router.POST("/signup", h.Signup)

			// Marshal `tc.body` into JSON so that it can be included in the HTTP request.
			body, err := json.Marshal(tc.body)

			// If the JSON marshaling fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/signup", bytes.NewReader(body))

			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, rec.Code)
			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
