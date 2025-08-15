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

func TestLoginHandler_Sucess(t *testing.T) {
	userRepo := NewInMemoryRepository()

	handler := NewHandler(userRepo)

	registrationData := `{"email":"test@example.com","password":"password123"}`
	reqregister, err := http.NewRequest("POST", "/register", bytes.NewReader([]byte(registrationData)))
	assert.NoError(t, err)

	rrRegister := httptest.NewRecorder()

	handler.RegisterUserHandler(rrRegister, reqregister)
	assert.Equal(t, http.StatusCreated, rrRegister.Code, "Expected status code 201 Created")

	// Agora, faça o login
	loginData := `{"email":"test@example.com","password":"password123"}`
	reqlogin, err := http.NewRequest("POST", "/login", bytes.NewReader([]byte(loginData)))
	assert.NoError(t, err)

	rrLogin := httptest.NewRecorder()
	handler.LoginUserHandler(rrLogin, reqlogin)

	assert.Equal(t, http.StatusOK, rrLogin.Code, "Expected status code 200 OK")
	assert.NotEmpty(t, rrLogin.Body.String(), "O corpo da resposta não deve estar vazio")
}
