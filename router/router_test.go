package router

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRouter_DNS tests the dns route which is /v1/dns
// It only tests the http response and not the location service
func TestRouter_DNS(t *testing.T) {
	router := Router()

	// Testing the handler when no data is provided
	t.Run("no-data", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/dns", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code, "should return 400 when there is no data")
	})

	// DNS request data that is not considered as a valid request
	invalidCases := [][]byte{
		[]byte(`{"x": 1, "y": 2.2, "z": 5, "vel": 10.22}`),
		[]byte(`{"x": "test", "y": "10.2", "z": "1.5", "vel": "20"}`),
		[]byte(`{"x": 4, "y": "6.7", "z": "11", "vel": "5.65"}`),
		[]byte(`{"x": "4", "y": "6.7", "z": 11, "vel": "5.65"}`),
		[]byte(`{"x": "21.43", "y": "2.2", "z": "5"}`),
		[]byte(`{"x": "0.43", "vel": "1.5"}`),
		[]byte(`{}`),
	}

	// All of the cases should cause the handler to return 400
	for i, c := range invalidCases {
		t.Run(fmt.Sprintf("with-invalid-data-%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			data := bytes.NewBuffer(c)
			req, _ := http.NewRequest("POST", "/v1/dns", data)
			router.ServeHTTP(w, req)
			assert.Equalf(t, http.StatusBadRequest, w.Code, "should return 400 with an invalid data: %s", c)
		})
	}

	// Valid DNS data
	validCases := [][]byte{
		[]byte(`{"x": "10", "y": "1.5", "z": "5", "vel": "10"}`),
		[]byte(`{"x": "15.65", "y": "4.23", "z": "5.32", "vel": "0.5"}`),
		[]byte(`{"x": "0", "y": "0", "z": "0", "vel": "0"}`),
		[]byte(`{"x": "-120", "y": "-14.5", "z": "10.15", "vel": "11.98"}`),
		[]byte(`{"x": "102.34", "y": "3.87", "z": "-98.1", "vel": "0"}`),
	}

	// All of the cases should cause the handler to return 200
	for i, c := range validCases {
		t.Run(fmt.Sprintf("with-valid-data-%d", i), func(t *testing.T) {
			w := httptest.NewRecorder()
			data := bytes.NewBuffer(c)
			req, _ := http.NewRequest("POST", "/v1/dns", data)
			router.ServeHTTP(w, req)
			assert.Equalf(t, http.StatusOK, w.Code, "should return 200 with a valid data: %s", c)
		})
	}
}
