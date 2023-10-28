package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"service-app/auth"
	"service-app/middlewares"
	"service-app/models"
	"service-app/models/mockmodels"
	"testing"
	"time"
)

func TestViewInventory(t *testing.T) {
	// Sets the Gin router mode to test.
	gin.SetMode(gin.TestMode)
	fakeClaims := jwt.RegisteredClaims{
		Subject: "1",
	}
	// MockUser struct initialization
	mockInventory := []models.Inventory{{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
			UpdatedAt: time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC),
		},
		ItemName:    "Black Shirt",
		Quantity:    10,
		Category:    "shirt",
		UserId:      1,
		CostPerItem: 6,
	}}
	// Define the list of test cases
	testCases := []struct {
		name              string                          // Name of the test case
		expectedStatus    int                             // Expected status of the response
		expectedResponse  string                          // Expected response body
		expectedInventory []models.Inventory              // Expected user after signup
		mockService       func(m *mockmodels.MockService) // Mock service function
	}{
		{
			name:             "OK",
			expectedStatus:   200,
			expectedResponse: `{"inv":[{"ID":1,"CreatedAt":"2006-01-01T01:01:01.000000001Z","UpdatedAt":"2006-01-01T01:01:01.000000001Z","DeletedAt":null,"item_name":"Black Shirt","quantity":10,"category":"shirt","user_id":1,"cost_per_item":6}],"total_cost":60}`,
			// Function for mocking service.
			// This simulates CreateInventory service and its return value.
			mockService: func(m *mockmodels.MockService) {

				var f float64 = 60
				m.EXPECT().ViewInventory(gomock.Any(), gomock.Any()).Times(1).
					Return(mockInventory, f, nil)
			},
		},
	}

	// Start a loop over `testCases` array where each element is represented by `tc`.
	for _, tc := range testCases {
		// Run a new test with the `tc.name` as its identifier.
		t.Run(tc.name, func(t *testing.T) {
			// Create a new Gomock controller.
			ctrl := gomock.NewController(t)

			// Create a mock Inventory using the Gomock controller.
			mockS := mockmodels.NewMockService(ctrl)

			// Apply the mock to the user service.
			tc.mockService(mockS)

			// Create a new instance of `models.Service` with the mock service.
			ms := models.NewStore(mockS)

			// Create a new context. This is typically passed between functions
			// carrying deadline, cancellation signals, and other request-scoped values.
			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, auth.Key, fakeClaims)
			// Create a fake TraceID. This would typically be used for request tracing.
			traceID := "fake-trace-id"
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middlewares.TraceIdKey, traceID)

			// Create a new Gin router.
			router := gin.New()

			// Create a new handler which uses the service model.
			h := handler{s: ms}

			// Register an endpoint and its handler with the router.
			router.POST("/view", h.ViewInventory)

			// Create a new HTTP POST request to "/signup".
			req, err := http.NewRequestWithContext(ctx, http.MethodPost, "/view", nil)
			// If the request creation fails, raise an error and stop the test.
			require.NoError(t, err)

			// Create a new HTTP Response Recorder. This is used to capture the HTTP response for analysis.
			resp := httptest.NewRecorder()

			// Pass the HTTP request to the router. This effectively "performs" the request and gets the associated response.
			router.ServeHTTP(resp, req)

			// Assert the returned HTTP status code is as expected.
			require.Equal(t, tc.expectedStatus, resp.Code)

			// Assert the response matches the expected response.
			require.Equal(t, tc.expectedResponse, string(resp.Body.Bytes()))
		})
	}
}
