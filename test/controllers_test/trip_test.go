package controllers_test

import (
	"fmt"
	"bytes"
    "net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func TestGetTrip200(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/trip/1"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetTrip404(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/trip/2"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostTrip201(t *testing.T) {
	userId := 1
	postPath := fmt.Sprintf("/trip/%v", userId)

	requestBody := fmt.Sprintf("{\"prefectureid\":\"%s\",\"title\":\"%s\",\"startdate\":\"%s\",\"enddate\":\"%s\",\"memo\":\"%s\",\"ispublic\":%t}", tripJson.PrefectureID, tripJson.Title, utils.TimeToString(tripJson.StartDate), utils.TimeToString(tripJson.EndDate), tripJson.Memo, tripJson.IsPublic)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostTrip400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"name1\":\"%s\"}", userJson.Name)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPutTrip201(t *testing.T) {
	newTitle := "茨城に行ってきた"
	requestBody := fmt.Sprintf("{\"prefectureid\":\"%s\",\"title\":\"%s\",\"startdate\":\"%s\",\"enddate\":\"%s\",\"memo\":\"%s\",\"ispublic\":%t}", tripJson.PrefectureID, newTitle, utils.TimeToString(tripJson.StartDate), utils.TimeToString(tripJson.EndDate), tripJson.Memo, tripJson.IsPublic)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	putPath := "/trip/2"
	req, _ := http.NewRequest("PUT", putPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPutTrip404(t *testing.T) {
	newTitle := "茨城に行ってきた"
	requestBody := fmt.Sprintf("{\"prefectureid\":\"%s\",\"title\":\"%s\",\"startdate\":\"%s\",\"enddate\":\"%s\",\"memo\":\"%s\",\"ispublic\":%t}", tripJson.PrefectureID, newTitle, utils.TimeToString(tripJson.StartDate), utils.TimeToString(tripJson.EndDate), tripJson.Memo, tripJson.IsPublic)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	putPath := "/trip/3"
	req, _ := http.NewRequest("PUT", putPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUploadTripImage201(t *testing.T) {
	for i := 0; i < 5; i ++ {
		requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
		requestJson := bytes.NewBufferString(requestBody)
		w := httptest.NewRecorder()
		postPath := "/tripimage/2"
		req, _ := http.NewRequest("POST", postPath, requestJson)
		Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}

func TestUploadTripImge400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
	requestJson := bytes.NewBufferString(requestBody)
	w := httptest.NewRecorder()
	postPath := "/tripimage/3"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteTripImage201(t *testing.T) {
	w := httptest.NewRecorder()
	deletePath := "/tripimage/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteTripImage404(t *testing.T) {
	w := httptest.NewRecorder()
	deletePath := "/tripimage/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteTrip201(t *testing.T) {
    w := httptest.NewRecorder()
	deletePath := "/trip/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteTrip404(t *testing.T) {
    w := httptest.NewRecorder()
	deletePath := "/trip/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

var tripJson = models.Trip{
	PrefectureID: "2",
	Title: "青森旅行",
	StartDate: utils.StringToTime("2024-02-23T08:43:42.709+09:00"),
	EndDate: utils.StringToTime("2024-02-25T08:43:42.709+09:00"),
	Memo: "りんごを食べました",
	IsPublic: true,
}