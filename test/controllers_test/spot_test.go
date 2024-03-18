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

func TestGetSpot200(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/spot/1"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSpot404(t *testing.T) {
    w := httptest.NewRecorder()
	getPath := "/spot/2"
	req, _ := http.NewRequest("GET", getPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPostSpot201(t *testing.T) {
	tripId := 1
	postPath := fmt.Sprintf("/spot/%v", tripId)

	requestBody := fmt.Sprintf("{\"tripid\":\"%v\",\"name\":\"%s\",\"date\":\"%s\",\"starttime\":\"%s\",\"endtime\":\"%s\",\"cost\":%v,\"memo\":\"%s\"}", tripId, spotJson.Name, utils.TimeToString(spotJson.Date), utils.TimeToString(spotJson.StartTime), utils.TimeToString(spotJson.EndTime), spotJson.Cost, spotJson.Memo)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostSpot400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"name1\":\"%s\"}", userJson.Name)
	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	postPath := "/user"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPutSpot201(t *testing.T) {
	newName := "茨城に行ってきた"
	tripId := 1

	requestBody := fmt.Sprintf("{\"tripid\":\"%v\",\"name\":\"%s\",\"date\":\"%s\",\"starttime\":\"%s\",\"endtime\":\"%s\",\"cost\":%v,\"memo\":\"%s\"}", tripId, newName, utils.TimeToString(spotJson.Date), utils.TimeToString(spotJson.StartTime), utils.TimeToString(spotJson.EndTime), spotJson.Cost, spotJson.Memo)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	putPath := "/spot/2"
	req, _ := http.NewRequest("PUT", putPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPutSpot404(t *testing.T) {
	newName := "茨城に行ってきた"
	tripId := 1
	requestBody := fmt.Sprintf("{\"tripid\":\"%v\",\"name\":\"%s\",\"date\":\"%s\",\"starttime\":\"%s\",\"endtime\":\"%s\",\"cost\":%v,\"memo\":\"%s\"}", tripId, newName, utils.TimeToString(spotJson.Date), utils.TimeToString(spotJson.StartTime), utils.TimeToString(spotJson.EndTime), spotJson.Cost, spotJson.Memo)

	requestJson := bytes.NewBufferString(requestBody)
    w := httptest.NewRecorder()
	putPath := "/spot/3"
	req, _ := http.NewRequest("PUT", putPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUploadSpotImage201(t *testing.T) {
	for i := 0; i < 5; i ++ {
		requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
		requestJson := bytes.NewBufferString(requestBody)
		w := httptest.NewRecorder()
		postPath := "/spotimage/2"
		req, _ := http.NewRequest("POST", postPath, requestJson)
		Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}
}

func TestUploadSpotImge400(t *testing.T) {
	requestBody := fmt.Sprintf("{\"encodeddata\":\"%s\"}", imageData)
	requestJson := bytes.NewBufferString(requestBody)
	w := httptest.NewRecorder()
	postPath := "/spotimage/3"
	req, _ := http.NewRequest("POST", postPath, requestJson)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteSpotImage201(t *testing.T) {
	w := httptest.NewRecorder()
	deletePath := "/spotimage/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteSpotImage404(t *testing.T) {
	w := httptest.NewRecorder()
	deletePath := "/spotimage/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteSpot201(t *testing.T) {
    w := httptest.NewRecorder()
	deletePath := "/spot/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteSpot404(t *testing.T) {
    w := httptest.NewRecorder()
	deletePath := "/spot/2"
	req, _ := http.NewRequest("DELETE", deletePath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
var spotJson = models.Spot{
	TripID: "1",
	Name: "袋田",
	StartTime: utils.StringToTime("2024-02-23T08:43:42.709+09:00"),
	EndTime: utils.StringToTime("2024-02-25T08:43:42.709+09:00"),
	Cost: 1000,
	Memo: "マイナスイオンがすごい",
}