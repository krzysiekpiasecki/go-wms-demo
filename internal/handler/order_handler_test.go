package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateOrderValidation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	orderHandler := &OrderHandler{}

	router.POST(
		"/orders",
		orderHandler.CreateOrder,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/orders",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(
		recorder,
		req,
	)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}
