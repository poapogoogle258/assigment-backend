package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPieFireDire(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		mockGetMeetText func() string
		expectedStatus  int
		expectedBody    map[string]int
	}{
		{
			name: "valid response",
			mockGetMeetText: func() string {
				return "Bacon ipsum dolor amet beef ribs chuck pork belly brisket."
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]int{
				"beef":    1,
				"ribs":    1,
				"chuck":   1,
				"pork":    1,
				"belly":   1,
				"brisket": 1,
			},
		},
		{
			name: "no meat words",
			mockGetMeetText: func() string {
				return "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
			},
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			service := NewPieFireDire(tt.mockGetMeetText)
			r.GET("/beef/summary", service.PieFireDire)

			req, _ := http.NewRequest(http.MethodGet, "/beef/summary", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var responseBody map[string]int
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}
