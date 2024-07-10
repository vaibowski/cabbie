package driver

import (
	"cabbie/driver/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpHandler(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        string
		mockService        func() *mocks.MockService
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "invalid request body",
			requestBody:        `invalid body`,
			mockService:        func() *mocks.MockService { return new(mocks.MockService) },
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "error unmarshalling request body\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/driver/signup", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()
			handler := SignUpHandler(tt.mockService())
			handler.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, tt.expectedStatusCode, resp.StatusCode)
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}
