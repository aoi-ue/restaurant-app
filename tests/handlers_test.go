package handlers

import (
	"net/http"
	"net/http/httptest"
	"restaurant-app/handlers"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func runRequest(r http.Handler, method, path string, body string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAddChoices(t *testing.T) {
	router := gin.Default()
	router.POST("/add-choices", handlers.AddChoices)

	body := `{"Name":"Restaurant A"}`
	req := runRequest(router, "POST", "/add-choices", body)
	assert.Equal(t, req.Code, http.StatusOK, "Response Should be 200 or OK")
}

func TestGetRandomChoices(t *testing.T) {
	router := gin.Default()
	router.GET("/get-random-choices", handlers.GetRandomChoices)

	body := ""
	req := runRequest(router, "GET", "/get-random-choices", body)

	assert.Equal(t, req.Code, http.StatusOK, "Response Should be 200 or OK")
}
