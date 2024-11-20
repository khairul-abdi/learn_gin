package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi TestPingRoute untuk menguji rute "/ping"
func TestPingRoute(t *testing.T) {
	// Memanggil fungsi setupRouter untuk membuat router
	router := router()

	// Membuat objek recorder HTTP baru
	w := httptest.NewRecorder()

	userLogin := Login{
		Username: "khairul",
		Password: "123456",
	}

	request, err := json.Marshal(userLogin)
	if err != nil {
		log.Fatal(err)
	}

	responseBody := bytes.NewBuffer(request)
	// Membuat HTTP request baru dengan method GET dan target rute "/ping"
	req, _ := http.NewRequest("POST", "/v1/login", responseBody)

	// Melakukan serve HTTP dengan router dan recorder HTTP
	router.ServeHTTP(w, req)

	// Memeriksa apakah kode respon HTTP sama dengan 200
	assert.Equal(t, http.StatusOK, w.Code)

	// Memeriksa apakah isi respon HTTP sama dengan `{"message":"pong"}`
	assert.Equal(t, `{"message":"success login"}`, w.Body.String())
}
