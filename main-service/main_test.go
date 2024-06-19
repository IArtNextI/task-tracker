package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHappyPathRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := NewMockDb()

	r := gin.Default()
	r.POST("/register", GetRegisterHandler(mockService))

	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString("{\"login\":\"login\", \"password\":\"password\"}"))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestDoubleRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := NewMockDb()

	r := gin.Default()
	r.POST("/register", GetRegisterHandler(mockService))

	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString("{\"login\":\"login\", \"password\":\"password\"}"))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "", w.Body.String())

	req, _ = http.NewRequest("POST", "/register", bytes.NewBufferString("{\"login\":\"login\", \"password\":\"password\"}"))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Equal(t, "", w.Body.String())
}
