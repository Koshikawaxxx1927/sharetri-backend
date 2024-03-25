package controllers_test

import (
	"fmt"
	"bytes"
    "net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"/app/models"
)

var userJson = models.User{
	Name: "千葉さん",
}

func TestGetUser200(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/user/1"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUser404(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/user/2"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostUser201(t *testing.T) {
	userJson := models.User{
		Name: "千葉さん",
	}
	requestBody := fmt.Sprintf("{\"name\":\"%s\"}", userJson.Name)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostUser400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"name1\":\"%s\"}", userJson.Name)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPutUser404(t *testing.T) {
	newName := "茨城さんだね"
	requestBody := fmt.Sprintf("{\"name\":\"%s\"}", newName)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user/2"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPutUser201(t *testing.T) {
	newName := "茨城さんだった"
	requestBody := fmt.Sprintf("{\"name\":\"%s\"}", newName)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user/2"
	req, _ := http.NewRequest("PUT", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUploadUserIcon201(t *testing.T) {
	for i := 0; i < 5; i ++ {
		requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
		requestJson := bytes.NewBufferString(requestBody)
		w := httptest.NewRecorder()
		postPath := "/usericon/2"
		req, _ := http.NewRequest("POST", postPath, requestJson)
		Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}

func TestUploadUserIcon400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
	requestJson := bytes.NewBufferString(requestBody)
	w := httptest.NewRecorder()
	postPath := "/usericon/3"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteUserIcon201(t *testing.T) {
	w := httptest.NewRecorder()
	postPath := "/usericon/2"
	req, _ := http.NewRequest("DELETE", postPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteUserIcon404(t *testing.T) {
	w := httptest.NewRecorder()
	postPath := "/usericon/2"
	req, _ := http.NewRequest("DELETE", postPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteUser201(t *testing.T) {
    w := httptest.NewRecorder()
	postPath := "/user/2"
	req, _ := http.NewRequest("DELETE", postPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteUser404(t *testing.T) {
    w := httptest.NewRecorder()
	postPath := "/user/2"
	req, _ := http.NewRequest("DELETE", postPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

var imageData = "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"