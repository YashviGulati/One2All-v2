package main

import (
	"gorilla/mux"
	"Projectv2/middlewares/basicauthmiddleware"
	"net/http"
	"Projectv2/handlecontrol"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func y() *mux.Router {
	y := mux.NewRouter()
	y.HandleFunc("/", handlecontrol.Home).Methods("GET")
	y.HandleFunc("/topics", handlecontrol.GetTopics).Methods("GET")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.GetSubByTopic))).Methods("GET")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.CreateTopic))).Methods("POST")
	y.HandleFunc("/topics/{name}/{sub}", handlecontrol.CreateSub).Methods("POST")
	y.Handle("/{name}/{msg}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.SendMsg))).Methods("POST")
	y.Handle("/topics/{name}", basicauthmiddleware.BasicAuthMiddleware(http.HandlerFunc(handlecontrol.DeleteTopicByName))).Methods("DELETE")

	return y
}

func TestHome(t *testing.T){
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)

}

func TestGetTopics(t *testing.T){
	request, _ := http.NewRequest("GET", "/topics", nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func TestGetSubByTopic(t *testing.T){
	request, _ := http.NewRequest("GET", "/topics/{name}",nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func TestCreateTopic(t *testing.T){
	request, _ := http.NewRequest("POST", "/topics/{name}",nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func TestDeleteTopicByName(t *testing.T){
	request, _ := http.NewRequest("DELETE", "/topics/{name}",nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

func TestSendMsg(t *testing.T){
	request, _ := http.NewRequest("POST", "/{name}/{msg}", nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Println(response.Body)
}

//func Create Sub


