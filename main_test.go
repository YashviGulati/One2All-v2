package main

import (
	"gorilla/mux"
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"fmt"
	"Projectv2/handlecontrol"
	"Projectv2/middlewares/basicauthmiddleware"
)

type Topic struct{
	name string
	sub string
	msg string
}

func y() *mux.Router {
	y := mux.NewRouter()
	y.HandleFunc("/", handlecontrol.Home).Methods("GET")
	y.HandleFunc("/topics", handlecontrol.GetTopics).Methods("GET")
	y.HandleFunc("/topics/{name}", handlecontrol.GetSubByTopic).Methods("GET")
	y.HandleFunc("/topics/{name}", handlecontrol.CreateTopic).Methods("POST")
	y.HandleFunc("/{name}/{msg}", handlecontrol.SendMsg).Methods("POST")
	y.HandleFunc("/topics/{name}/{sub}", handlecontrol.CreateSub).Methods("POST")
	//y.HandleFunc("/topics/{name}", handlecontrol.DeleteTopicByName).Methods("DELETE")
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
	topicname:="topic03"
	request, _ := http.NewRequest("GET", "/topics/"+topicname, nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Print(response.Body)
}

func TestCreateTopic(t *testing.T){
	topicname:="topic04"
	request, _ := http.NewRequest("POST", "/topics/"+topicname,nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Print(response.Body)
}

func TestDeleteTopicByName(t *testing.T){
	topicname:="topic04"
	request, _ := http.NewRequest("DELETE", "/topics/"+topicname, nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Access denied")
}

func TestSendMsg(t *testing.T){
	topicname:="topic03"
	message:="Hi"
	request, _ := http.NewRequest("POST", "/"+topicname+"/"+message,nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Print(response.Body)
}

func TestCreateSub(t *testing.T){
	topicname:="topic01"
	subname:="yashvi.gulati@gmail.com"
	request, _ := http.NewRequest("POST", "/topics/"+topicname+"/"+subname,nil)
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	fmt.Print(response.Body)
}

/*
Documentation:

func NewRequest():
	NewRequest returns a new Request given a method, URL, and optional body

func NewRecorder():
	NewRecorder returns an initialized ResponseRecorder.

func ServeHTTP():
	ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.

func NewBuffer():
	creates a buffer to read an existing data

func Marshal():
	Marshal returns the JSON encoding
 */