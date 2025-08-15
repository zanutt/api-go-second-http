package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {

	userRepo := NewInMemoryRepository()

	handler := NewHandler(userRepo)

	registrationData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "test@example.com",
		Password: "password123",
	}

	body, err := json.Marshal(registrationData)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/register", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	handler.RegisterUserHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code, "Expected status code 201 Created")
	assert.NotEmpty(t, rr.Body.String(), "O corpo da resposta não deve estar vazio")

}
