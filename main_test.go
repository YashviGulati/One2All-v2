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

	"encoding/json"
	"bytes"
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

	topic:=&Topic{
		name:"topic03",
	}
	jsontopic,_:=json.Marshal(topic)
	request, _ := http.NewRequest("GET", "/topics/{name}", bytes.NewBuffer(jsontopic))
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Admin access needed")
}

func TestCreateTopic(t *testing.T){
	topic:=&Topic{
		name:"topic04",
	}
	jsontopic,_:=json.Marshal(topic)
	request, _ := http.NewRequest("POST", "/topics/{name}",bytes.NewBuffer(jsontopic))
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Admin access needed")
}

func TestDeleteTopicByName(t *testing.T){
	topic:=&Topic{
		name:"topic04",
	}
	jsontopic,_:=json.Marshal(topic)
	request, _ := http.NewRequest("DELETE", "/topics/{name}", bytes.NewBuffer(jsontopic))
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Admin access needed")
}

func TestSendMsg(t *testing.T){
	topic:=&Topic{
		name:"topic04",
		msg:"Hello",
	}
	jsontopic,_:=json.Marshal(topic)
	request, _ := http.NewRequest("POST", "/{name}/{msg}",bytes.NewBuffer(jsontopic))
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Admin access needed")
}

func TestCreateSub(t *testing.T){
	topic:=&Topic{
		name:"topic04",
		sub:"gulati.yashvi@gmail.com",
	}
	jsontopic,_:=json.Marshal(topic)
	request, _ := http.NewRequest("POST", "/topics/{name}",bytes.NewBuffer(jsontopic))
	response := httptest.NewRecorder()
	y().ServeHTTP(response, request)
	assert.Equal(t, 401, response.Code, "OK response is expected")
	fmt.Print(response.Body)
	fmt.Println("Admin access needed")
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